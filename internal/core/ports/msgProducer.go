package ports

import "github.com/TomasCruz/users/internal/core/entities"

type MsgProducer interface {
	Close()
	PublishUserModification(user entities.User, modificationType entities.UserModification) error
}
