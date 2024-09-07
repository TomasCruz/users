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

func (c Core) ListUser(filter map[string][]string, pageSize, pageNumber int) ([]entities.User, int64, error) {
	users, totalCount, err := c.db.ListUser(filter, pageSize, pageNumber)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
