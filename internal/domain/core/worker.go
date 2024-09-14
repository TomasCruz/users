package core

import (
	"fmt"

	"github.com/TomasCruz/users/internal/domain/entities"
)

func (cr Core) ConsumeUserCreatedMsg(user entities.User) error {
	// this is a dummy, so what
	cr.logger.Info(nil, fmt.Sprintf("User %s got created!!!", user.UserID.String()))
	return nil
}
