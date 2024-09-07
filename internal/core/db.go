package core

import (
	"io"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/google/uuid"
)

// DB is an interface through which to talk with DB
type DB interface {
	io.Closer
	Health(dbString string) error
	// CreateUser(userID uuid.UUID, firstName, lastName, nickName, pswdHash, email, country string, createdAt, updatedAt time.Time) (entities.User, error)
	GetUserByID(userID uuid.UUID) (entities.User, error)
	ResultCountQuery(filteredQuery string, args []interface{}) (int64, error)
	QueriesAndParameterNamesForListUser() (string, string, map[string]string)
	// ListUser(parametrizedQueryString string, args []interface{}) ([]entities.User, error)
	ListUser(filter map[string][]string, pageSize, pageNumber int) ([]entities.User, int64, error)
	// UpdateUser(userID uuid.UUID, updatedAt time.Time, req entities.UpdateUserReq) (entities.User, error)
	// DeleteUser(userID uuid.UUID) (entities.User, error)
}
