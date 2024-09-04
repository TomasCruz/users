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
func (c Core) ListUser(filter entities.Filter, paginator entities.Paginator) ([]entities.UserResp, int64, error) {
	users, totalCount, err := c.db.ListUser(filter, paginator)
	if err != nil {
		return nil, 0, err
	}

	l := len(users)
	resps := make([]entities.UserResp, 0, l)
	for _, u := range users {
		resps = append(resps, entities.UserRespFromUser(u))
	}

	return resps, totalCount, nil
}
