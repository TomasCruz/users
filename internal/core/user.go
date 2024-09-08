package core

import (
	"time"

	"github.com/TomasCruz/users/internal/entities"
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

func (c Core) ListUser(filter map[string][]string, pageSize, pageNumber int) ([]entities.User, int64, error) {
	users, totalCount, err := c.db.ListUser(filter, pageSize, pageNumber)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (c Core) CreateUser(req entities.CreateUserReq) (entities.UserResp, error) {
	userID := uuid.New()
	now := time.Now().UTC()

	user, err := c.db.CreateUser(userID, req.FirstName, req.LastName, req.PswdHash, req.Email, req.Country, now, now)
	if err != nil {
		if errors.Is(err, entities.ErrExistingEmail) {
			return entities.UserResp{}, err
		}

		return entities.UserResp{}, err
	}

	resp := entities.UserRespFromUser(user)
	err = c.msg.PublishUserModification(resp, c.config.CreateUserTopic)
	if err != nil {
		return entities.UserResp{}, err
	}

	return resp, nil
}
