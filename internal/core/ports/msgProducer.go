package ports

import (
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
)

type MsgProducer interface {
	SendUserMsg(user entities.User, modificationType entities.UserModification, timeout time.Duration) ([]byte, error)
	Close() error
}
