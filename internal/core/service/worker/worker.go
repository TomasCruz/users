package worker

import (
	"fmt"

	"github.com/TomasCruz/users/internal/core/entities"
)

func (svc WorkerUserService) ConsumeUserCreatedMsg(user entities.User) error {
	// this is a dummy, so what
	svc.logger.Info(nil, fmt.Sprintf("User %s got created!!!", user.UserID.String()))
	return nil
}
