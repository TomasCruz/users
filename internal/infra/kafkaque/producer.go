package kafkaque

import (
	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/pkg/errors"
)

func (k kafkaProducer) PublishUserEvent(user entities.User, modificationType entities.UserModification) error {
	switch modificationType {
	case entities.CREATE_MODIFICATION:
		return k.publishUserCreatedEvent(user)
	case entities.UPDATE_MODIFICATION:
		return nil
	case entities.DELETE_MODIFICATION:
		return nil
	default:
		return errors.WithStack(entities.ErrBadMsgType)
	}
}
