package core

import (
	"time"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (cr Core) GetUserByID(userID uuid.UUID) (entities.User, error) {
	user, err := cr.db.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (cr Core) ListUser(filter entities.UserFilter, pageSize, pageNumber int) ([]entities.User, int64, error) {
	users, totalCount, err := cr.db.ListUser(filter, pageSize, pageNumber)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (cr Core) CreateUser(req entities.UserDTO) (entities.User, error) {
	if _, err := cr.db.GetUserByEmail(*req.Email); err != nil {
		// ignore no rows error
		if !errors.Is(err, entities.ErrNonexistingUser) {
			return entities.User{}, err
		}
	} else {
		// user found by email
		return entities.User{}, entities.ErrExistingEmail
	}

	userID := uuid.New()
	now := time.Now().UTC()

	user, err := cr.db.CreateUser(req, userID, now, now)
	if err != nil {
		return entities.User{}, err
	}

	err = cr.msgProducer.PublishUserModification(user, entities.CREATE_MODIFICATION)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
