//go:build unit
// +build unit

package app

import (
	"errors"
	"testing"
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/TomasCruz/users/tests"
	"github.com/TomasCruz/users/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateUser(t *testing.T) {
	userID := uuid.New()
	firstName := "firstName"
	lastName := "lastName"
	hash := "***"
	email := "a@a.com"
	country := "USA"

	now := time.Now()
	req := entities.UserDTO{
		UserID:    &userID,
		FirstName: &firstName,
		LastName:  &lastName,
		PswdHash:  &hash,
		Email:     &email,
		Country:   &country,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	user := entities.User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		PswdHash:  hash,
		Email:     email,
		Country:   country,
		CreatedAt: now,
		UpdatedAt: now,
	}

	testCases := []struct {
		name string

		req entities.UserDTO

		getUserByEmailErr error

		sholudCreateUser bool
		createdUser      entities.User
		createUserErr    error

		shouldPublish bool
		publishErr    error

		user entities.User
		err  error
	}{
		{
			name: "success",

			req: req,

			getUserByEmailErr: entities.ErrNonexistingUser,

			sholudCreateUser: true,
			createdUser:      user,
			createUserErr:    nil,

			shouldPublish: true,
			publishErr:    nil,

			user: user,
			err:  nil,
		},
		{
			name: "publish error",

			req: req,

			getUserByEmailErr: entities.ErrNonexistingUser,

			sholudCreateUser: true,
			createdUser:      user,
			createUserErr:    nil,

			shouldPublish: true,
			publishErr:    errors.New("boom"),

			user: entities.User{},
			err:  errors.New("boom"),
		},
		{
			name: "db.CreateUser error",

			req: req,

			getUserByEmailErr: entities.ErrNonexistingUser,

			sholudCreateUser: true,
			createdUser:      user,
			createUserErr:    errors.New("boom"),

			user: entities.User{},
			err:  errors.New("boom"),
		},
		{
			name: "GetUserByEmail no error",

			req: req,

			getUserByEmailErr: nil,

			user: entities.User{},
			err:  entities.ErrExistingEmail,
		},
		{
			name: "GetUserByEmail general error",

			req: req,

			getUserByEmailErr: errors.New("boom"),

			user: entities.User{},
			err:  errors.New("boom"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db := &mocks.DB{}
			msg := &mocks.MsgProducer{}
			logger := &mocks.Logger{}

			db.On("GetUserByEmail", *tt.req.Email).
				Return(entities.User{}, tt.getUserByEmailErr)

			if tt.sholudCreateUser {
				db.On("CreateUser", req, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).
					Return(tt.createdUser, tt.createUserErr)
			}

			if tt.shouldPublish {
				msg.On("PublishUserModification", tt.createdUser, entities.CREATE_MODIFICATION).
					Return(tt.publishErr)
			}

			svc := NewAppUserService(db, msg, logger)
			resultUser, err := svc.CreateUser(tt.req)

			tests.AssertEqualError(t, tt.err, err, "should return expected error")
			assert.Equal(t, tt.user, resultUser, "should return expected user FirstName")
		})
	}
}
