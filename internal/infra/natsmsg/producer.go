package natsmsg

import (
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
)

func (nProd natsProducer) SendUserMsg(user entities.User, modificationType entities.UserModification, timeout time.Duration) ([]byte, error) {
	switch modificationType {
	case entities.CREATE_MODIFICATION:
		return nProd.sendUserCreatedMsg(user, timeout)
	case entities.UPDATE_MODIFICATION:
		return nil, nil
	case entities.DELETE_MODIFICATION:
		return nil, nil
	}

	return nil, nil
}
