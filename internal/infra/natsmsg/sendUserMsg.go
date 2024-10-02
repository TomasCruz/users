package natsmsg

import (
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/pkg/errors"
)

func (nProd natsProducer) sendUserCreatedMsg(user entities.User, timeout time.Duration) ([]byte, error) {
	msg, err := nProd.nc.Request(nProd.config.NatsSubjectCreateUser, []byte(user.UserID.String()), timeout)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return msg.Data, nil
}
