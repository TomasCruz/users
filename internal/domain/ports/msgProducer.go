package ports

import "github.com/TomasCruz/users/internal/domain/entities"

type MsgProducer interface {
	Close()
	PublishUserModification(user entities.User, modificationType entities.UserModification) error
}
