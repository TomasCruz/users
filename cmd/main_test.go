package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"
	"testing"

	"github.com/TomasCruz/users/internal/app"
	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/entities"
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
	a := app.App{
		EnvFile:     "../.env.test",
		ServerReady: serverReady,
	}

	ts.config, _ = app.ConfigFromEnvVars(a.EnvFile)

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
	healthRespStatus, _ := ts.doReq("GET", fmt.Sprintf("http://localhost:%s/health", ts.config.Port), nil, nil, "", "")
	ts.Equal(http.StatusNoContent, healthRespStatus, "HealthResp status unexpected")
}

func (ts *testSuite) TestCreateUser() {
	// create user
	createUserHeader := http.Header{}
	createUserHeader.Add("Content-Type", "application/json")

	createUserReq := entities.CreateUserReq{
		FirstName: "FirstName",
		LastName:  "LastName",
		PswdHash:  "$2a$10$wrfLakZMCZHQStxyvmfmWuIF8ovj2Tcbdc9tH3VEf8MPWntBLg55W",
		Email:     "a@a.com",
		Country:   "UK",
	}

	createUserReqBytes, err := json.Marshal(createUserReq)
	ts.NoError(err, "CreateUserReq Marshal error")

	createUserReqBuffer := bytes.NewBuffer(createUserReqBytes)
	createUserRespStatus, createUserRespBodyBytes := ts.doReq("PUT", fmt.Sprintf("http://localhost:%s/users", ts.config.Port), createUserReqBuffer, createUserHeader, "CreateUserReq", "CreateUserResp")
	ts.Equal(http.StatusCreated, createUserRespStatus, "CreateUserResp status unexpected")

	var createUserResp entities.UserResp
	err = json.Unmarshal(createUserRespBodyBytes, &createUserResp)
	ts.NoError(err, "CreateUserResp Unmarshal error")

	// get user
	getUserRespStatus, getUserRespBodyBytes := ts.doReq("GET", fmt.Sprintf("http://localhost:%s/users/%s", ts.config.Port, createUserResp.UserID), nil, nil, "", "")
	ts.Equal(http.StatusOK, getUserRespStatus, "GetUser status unexpected")

	var getUserResp entities.UserResp
	err = json.Unmarshal(getUserRespBodyBytes, &getUserResp)
	ts.NoError(err, "GetUserResp Unmarshal error")
	ts.Equal(createUserReq.FirstName, getUserResp.FirstName, "GetUserResp FirstName unexpected")
	ts.Equal(createUserReq.LastName, getUserResp.LastName, "GetUserResp LastName unexpected")
	ts.Equal(createUserReq.PswdHash, getUserResp.PswdHash, "GetUserResp PswdHash unexpected")
	ts.Equal(createUserReq.Email, getUserResp.Email, "GetUserResp Email unexpected")
	ts.Equal(createUserReq.Country, getUserResp.Country, "GetUserResp Country unexpected")
}

func (ts *testSuite) doReq(method, url string, reqBuffer *bytes.Buffer, headers http.Header, reqStructName, respStructName string) (status int, respBytes []byte) {
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

	status = httpResp.StatusCode
	respBytes, err = io.ReadAll(httpResp.Body)
	ts.NoError(err, fmt.Sprintf("%s ReadAll error", respStructName))

	return
}
