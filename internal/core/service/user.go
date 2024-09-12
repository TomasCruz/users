package core

import (
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (svc Service) GetUserByID(userID uuid.UUID) (entities.User, error) {
	user, err := svc.db.GetUserByID(userID)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (svc Service) ListUser(filter entities.UserFilter, pageSize, pageNumber int) ([]entities.User, int64, error) {
	users, totalCount, err := svc.db.ListUser(filter, pageSize, pageNumber)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (svc Service) CreateUser(req entities.UserDTO) (entities.User, error) {
	if _, err := svc.db.GetUserByEmail(*req.Email); err != nil {
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

	user, err := svc.db.CreateUser(userID, *req.FirstName, *req.LastName, *req.PswdHash, *req.Email, *req.Country, now, now)
	if err != nil {
		return entities.User{}, err
	}

	err = svc.msg.PublishUserModification(user, entities.CREATE_MODIFICATION)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
