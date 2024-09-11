package core

import (
	"time"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (c Core) GetUserByID(userID uuid.UUID) (entities.User, error) {
	user, err := c.db.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (c Core) ListUser(filter map[string]map[string]struct{}, pageSize, pageNumber int) ([]entities.User, int64, error) {
	users, totalCount, err := c.db.ListUser(filter, pageSize, pageNumber)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (c Core) CreateUser(req entities.UserDTO) (entities.User, error) {
	if _, err := c.db.GetUserByEmail(*req.Email); err != nil {
		// ignore no rows error
		if !errors.Is(err, entities.ErrNonexistingUser) {
			return entities.User{}, err
		}
	} else {
		// no error, user found by email
		return entities.User{}, entities.ErrExistingEmail
	}

	userID := uuid.New()
	now := time.Now().UTC()

	user, err := c.db.CreateUser(userID, *req.FirstName, *req.LastName, *req.PswdHash, *req.Email, *req.Country, now, now)
	if err != nil {
		return entities.User{}, err
	}

	err = c.msg.PublishUserModification(user, entities.CREATE_MODIFICATION)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
