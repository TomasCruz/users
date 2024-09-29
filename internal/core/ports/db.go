package ports

import (
	"io"
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/google/uuid"
)

type DB interface {
	io.Closer
	Health() error
	CreateUser(req entities.UserDTO, userID uuid.UUID, createdAt, updatedAt time.Time) (entities.User, error)
	GetUserByID(userID uuid.UUID) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	ListUser(filter entities.UserFilter, pageSize, pageNumber int) ([]entities.User, int64, error)
	// UpdateUser(userID uuid.UUID, updatedAt time.Time, req UpdateUserReq) (User, error)
	// DeleteUser(userID uuid.UUID) (User, error)
}
