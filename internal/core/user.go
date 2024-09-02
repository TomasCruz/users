package core

import (
	"github.com/TomasCruz/users/internal/entities"
	"github.com/google/uuid"
)

func (c Core) GetUserByID(userID uuid.UUID) (entities.User, error) {
	user, err := c.db.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
