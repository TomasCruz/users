package ports

import "github.com/TomasCruz/users/internal/domain/entities"

type Msg interface {
	Close()
	PublishUserModification(user entities.User, modificationType entities.UserModification) error
}
