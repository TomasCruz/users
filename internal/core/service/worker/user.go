package worker

import (
	"fmt"
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
)

func (svc WorkerUserService) ConsumeUserCreatedEvent(user entities.User) error {
	// this is a dummy, so what
	svc.logger.Info(nil, fmt.Sprintf("User %s got created!!!", user.UserID.String()))

	// another dummy
	rep, err := svc.msgProducer.SendUserMsg(user, entities.CREATE_MODIFICATION, 1*time.Second)
	if err != nil {
		return err
	}

	svc.logger.Info(nil, fmt.Sprintf("NATS create user resp: %s", string(rep)))
	return nil
}
