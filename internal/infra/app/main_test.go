//go:build integration
// +build integration

package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"testing"

	"github.com/TomasCruz/users/internal/handlers/httphandler"
	"github.com/TomasCruz/users/internal/infra/configuration"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	config configuration.Config
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &testSuite{})
}

func (ts *testSuite) SetupSuite() {
	serverReady := make(chan struct{})
	a := App{
		EnvFile:     "../../../.env.test",
		ServerReady: serverReady,
	}

	ts.config, _ = configuration.ConfigFromEnvVars(a.EnvFile)

	go a.Start()
	<-serverReady
}

func (ts *testSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (ts *testSuite) SetupTest() {
}

func (ts *testSuite) TearDownTest() {
}

func (ts *testSuite) TestHealth() {
	healthRespStatus, _, _ := ts.doReq("GET", fmt.Sprintf("http://localhost:%s/health", ts.config.Port), nil, nil, "", "")
	ts.Equal(http.StatusNoContent, healthRespStatus, "HealthResp status unexpected")
}

func (ts *testSuite) TestCreateUser() {
	// create user
	createUserHeader := http.Header{}
	createUserHeader.Add("Content-Type", "application/json")

	createUserReq := httphandler.CreateUserReq{
		FirstName: "FirstName",
		LastName:  "LastName",
		PswdHash:  "$2a$10$wrfLakZMCZHQStxyvmfmWuIF8ovj2Tcbdc9tH3VEf8MPWntBLg55W",
		Email:     "a@a.com",
		Country:   "MNG",
	}

	createUserReqBytes, err := json.Marshal(createUserReq)
	ts.NoError(err, "CreateUserReq Marshal error")

	createUserReqBuffer := bytes.NewBuffer(createUserReqBytes)
	createUserRespStatus, createUserRespBodyBytes, _ := ts.doReq("PUT", fmt.Sprintf("http://localhost:%s/users", ts.config.Port), createUserReqBuffer, createUserHeader, "CreateUserReq", "CreateUserResp")
	ts.Equal(http.StatusCreated, createUserRespStatus, "CreateUserResp status unexpected")

	var createUserResp httphandler.UserResp
	err = json.Unmarshal(createUserRespBodyBytes, &createUserResp)
	ts.NoError(err, "CreateUserResp Unmarshal error")

	// get user
	getUserRespStatus, getUserRespBodyBytes, _ := ts.doReq("GET", fmt.Sprintf("http://localhost:%s/users/%s", ts.config.Port, createUserResp.UserID), nil, nil, "", "")
	ts.Equal(http.StatusOK, getUserRespStatus, "GetUser status unexpected")

	var getUserResp httphandler.UserResp
	err = json.Unmarshal(getUserRespBodyBytes, &getUserResp)
	ts.NoError(err, "GetUserResp Unmarshal error")
	ts.Equal(createUserReq.FirstName, getUserResp.FirstName, "GetUserResp FirstName unexpected")
	ts.Equal(createUserReq.LastName, getUserResp.LastName, "GetUserResp LastName unexpected")
	ts.Equal(createUserReq.PswdHash, getUserResp.PswdHash, "GetUserResp PswdHash unexpected")
	ts.Equal(createUserReq.Email, getUserResp.Email, "GetUserResp Email unexpected")
	ts.Equal(createUserReq.Country, getUserResp.Country, "GetUserResp Country unexpected")
}

func (ts *testSuite) TestListUser() {
	// docker cp ../tests/scripts/users.sql pgdb_test:/docker-entrypoint-initdb.d/users.sql
	cmdCopy := exec.Command("docker", "cp", "../../../tests/scripts/users.sql", "pgdb_test:/docker-entrypoint-initdb.d/users.sql")
	_, err := cmdCopy.Output()
	if err != nil {
		ts.FailNow("docker Output failed", err.Error())
	}

	// docker exec -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test pgdb_test psql test test -f docker-entrypoint-initdb.d/users.sql
	cmdScript := exec.Command("docker", "exec", "-e", "POSTGRES_USER=test", "-e", "POSTGRES_PASSWORD=test", "pgdb_test", "psql", "test", "test", "-f", "docker-entrypoint-initdb.d/users.sql")
	_, err = cmdScript.Output()
	if err != nil {
		ts.FailNow("docker Output failed", err.Error())
	}

	indUUID, _ := uuid.Parse("b0be889c-147e-40d1-8e3c-d116990b2e74")
	firstNameIND := "FirstName12"
	tests := []struct {
		name            string
		url             string
		listCondition   func(int) bool
		totalCondition  func(int) bool
		resultCondition func(int) bool
		id              *uuid.UUID
		firstName       *string
	}{
		{
			name:            "full list user",
			url:             fmt.Sprintf("http://localhost:%s/users", ts.config.Port),
			listCondition:   func(x int) bool { return x >= 14 },
			totalCondition:  func(x int) bool { return x >= 14 },
			resultCondition: func(x int) bool { return x >= 14 },
		},
		{
			name:            "IND list user",
			url:             fmt.Sprintf("http://localhost:%s/users?country=IND", ts.config.Port),
			listCondition:   func(x int) bool { return x == 1 },
			totalCondition:  func(x int) bool { return x == 1 },
			resultCondition: func(x int) bool { return x == 1 },
			id:              &indUUID,
			firstName:       &firstNameIND,
		},
		{
			name:            "SRB list user 1",
			url:             fmt.Sprintf("http://localhost:%s/users?country=SRB&page-number=1&page-size=2", ts.config.Port),
			listCondition:   func(x int) bool { return x == 2 },
			totalCondition:  func(x int) bool { return x == 3 },
			resultCondition: func(x int) bool { return x == 2 },
		},
		{
			name:            "SRB list user 2",
			url:             fmt.Sprintf("http://localhost:%s/users?country=SRB&page-number=2&page-size=2", ts.config.Port),
			listCondition:   func(x int) bool { return x == 1 },
			totalCondition:  func(x int) bool { return x == 3 },
			resultCondition: func(x int) bool { return x == 1 },
		},
		{
			name:            "SRB list user 3",
			url:             fmt.Sprintf("http://localhost:%s/users?country=SRB&page-number=1&page-size=3", ts.config.Port),
			listCondition:   func(x int) bool { return x == 3 },
			totalCondition:  func(x int) bool { return x == 3 },
			resultCondition: func(x int) bool { return x == 3 },
		},
		{
			name:            "SRB list user 4",
			url:             fmt.Sprintf("http://localhost:%s/users?country=SRB&page-number=2&page-size=3", ts.config.Port),
			listCondition:   func(x int) bool { return x == 0 },
			totalCondition:  func(x int) bool { return x == 3 },
			resultCondition: func(x int) bool { return x == 0 },
		},
	}

	for _, tt := range tests {
		listUserRespStatus, listUserRespBodyBytes, listUserRespHeader := ts.doReq("GET", tt.url, nil, nil, "", "")
		ts.Equal(http.StatusOK, listUserRespStatus, fmt.Sprintf("%s: GetUser status unexpected", tt.name))

		var listUserResp []httphandler.UserResp
		err = json.Unmarshal(listUserRespBodyBytes, &listUserResp)
		ts.NoError(err, fmt.Sprintf("%s: list UserResp Unmarshal error: ", tt.name))

		ts.Condition(func() bool {
			return tt.listCondition(len(listUserResp))
		}, fmt.Sprintf("%s: listUserResp length unexpected", tt.name))
		ts.Condition(func() bool {
			return ts.headerFunc("X-Total-Count", listUserRespHeader, tt.totalCondition)
		}, fmt.Sprintf("%s: listUserResp X-Total-Count unexpected", tt.name))
		ts.Condition(func() bool {
			return ts.headerFunc("X-Result-Count", listUserRespHeader, tt.resultCondition)
		}, fmt.Sprintf("%s: listUserResp X-Result-Count unexpected", tt.name))

		if tt.id != nil {
			ts.Equal(*tt.id, listUserResp[0].UserID, fmt.Sprintf("%s: listUserResp ID unexpected", tt.name))
		}
		if tt.firstName != nil {
			ts.Equal(*tt.firstName, listUserResp[0].FirstName, fmt.Sprintf("%s: listUserResp FirstName unexpected", tt.name))
		}
	}
}

func (ts *testSuite) headerFunc(hdrName string, respHeader map[string][]string, f func(int) bool) bool {
	if respHeader == nil {
		return false
	}

	xStrings, present := respHeader[hdrName]
	if !present || len(xStrings) == 0 {
		return false
	}

	x, err := strconv.Atoi(xStrings[0])
	if err != nil {
		return false
	}

	return f(x)
}

func (ts *testSuite) doReq(method, url string, reqBuffer *bytes.Buffer, headers http.Header, reqStructName, respStructName string) (int, []byte, map[string][]string) {
	if reqBuffer == nil {
		reqBuffer = bytes.NewBuffer([]byte{})
	}

	httpReq, err := http.NewRequest(method, url, reqBuffer)
	ts.NoError(err, fmt.Sprintf("%s NewRequest error", reqStructName))

	if len(headers) != 0 {
		httpReq.Header = headers
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	ts.NoError(err, fmt.Sprintf("%s Do error", reqStructName))
	defer httpResp.Body.Close()

	status := httpResp.StatusCode
	respBytes, err := io.ReadAll(httpResp.Body)
	ts.NoError(err, fmt.Sprintf("%s ReadAll error", respStructName))

	respHeader := httpResp.Header

	return status, respBytes, respHeader
}
