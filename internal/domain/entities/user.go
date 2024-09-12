package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID
	FirstName string
	LastName  string
	PswdHash  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
