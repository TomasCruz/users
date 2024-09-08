package core

import (
	"io"
	"time"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/google/uuid"
)

// DB is an interface through which to talk with DB
type DB interface {
	io.Closer
	Health() error
	CreateUser(userID uuid.UUID, firstName, lastName, pswdHash, email, country string, createdAt, updatedAt time.Time) (entities.User, error)
	GetUserByID(userID uuid.UUID) (entities.User, error)
	ListUser(filter map[string][]string, pageSize, pageNumber int) ([]entities.User, int64, error)
	// UpdateUser(userID uuid.UUID, updatedAt time.Time, req entities.UpdateUserReq) (entities.User, error)
	// DeleteUser(userID uuid.UUID) (entities.User, error)
}
