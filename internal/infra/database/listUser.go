package database

import (
	"database/sql"
	"time"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (pDB postgresDB) queriesAndParameterNamesForListUser() (string, string, map[string]string) {
	basicQuery := `SELECT user_id, first_name, last_name, pswd_hash, email, country, created_at, updated_at FROM users`
	parameterNames := map[string]string{"country": "country"}

	return basicQuery, "", parameterNames
}

func (pDB postgresDB) ListUser(filter entities.UserFilter, pageSize, pageNumber int) ([]entities.User, int64, error) {
	basicQuery, orderByQuery, parameterNames := pDB.queriesAndParameterNamesForListUser()
	filteredQuery, args := buildFilteredQuery(basicQuery, filter, parameterNames)

	totalCount, err := pDB.resultCountQuery(filteredQuery, args)
	if err != nil {
		return nil, 0, err
	}

	parametrizedQuery, limit, offset := buildPaginatedQuery(filteredQuery, orderByQuery, pageSize, pageNumber, len(args))
	if limit != 0 {
		args = append(args, limit, offset)
	}

	var rows *sql.Rows
	if rows, err = pDB.db.Query(parametrizedQuery, args...); err != nil {
		return nil, 0, errors.Wrap(entities.ErrListUser, err.Error())
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var (
			userID    uuid.UUID
			firstName string
			lastName  string
			pswdHash  string
			email     string
			country   string
			createdAt time.Time
			updatedAt time.Time
		)

		if err = rows.Scan(&userID, &firstName, &lastName, &pswdHash, &email, &country, &createdAt, &updatedAt); err != nil {
			return nil, 0, errors.Wrap(entities.ErrListUser, err.Error())
		}

		user := entities.User{
			UserID:    userID,
			FirstName: firstName,
			LastName:  lastName,
			PswdHash:  pswdHash,
			Email:     email,
			Country:   country,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		users = append(users, user)
	}

	return users, totalCount, nil
}
