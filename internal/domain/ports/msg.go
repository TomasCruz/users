package ports

import "github.com/TomasCruz/users/internal/domain/entities"

// Msg is an interface through which to talk with DB
type Msg interface {
	Close()
	PublishUserModification(user entities.User, modificationType entities.UserModification) error
}
