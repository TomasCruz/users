package ports

import "github.com/TomasCruz/users/internal/core/entities"

type QueueProducer interface {
	Close()
	PublishUserEvent(user entities.User, modificationType entities.UserModification) error
}
