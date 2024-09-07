package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/TomasCruz/users/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (pDB postgresDB) QueriesAndParameterNamesForListUser() (string, string, map[string]string) {
	basicQuery := `SELECT user_id, first_name, last_name, pswd_hash, email, country, created_at, updated_at FROM users`
	parameterNames := map[string]string{"country": "country"}

	return basicQuery, "", parameterNames
}

func (pDB postgresDB) ResultCountQuery(filteredQuery string, args []interface{}) (int64, error) {
	countFilteredQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS a", filteredQuery)

	var totalCount int64
	err := pDB.db.QueryRow(countFilteredQuery, args...).Scan(&totalCount)
	if err != nil {
		return 0, errors.Wrap(entities.ErrCountFilteredQuery, err.Error())
	}

	return totalCount, nil
}

func (pDB postgresDB) ListUser(filter map[string][]string, pageSize, pageNumber int) ([]entities.User, int64, error) {
	basicQuery, orderByQuery, parameterNames := pDB.QueriesAndParameterNamesForListUser()
	filteredQuery, args := utils.BuildFilteredPostgresQuery(basicQuery, filter, parameterNames)

	totalCount, err := pDB.ResultCountQuery(filteredQuery, args)
	if err != nil {
		return nil, 0, err
	}

	parametrizedQuery, limit, offset := utils.BuildPaginatedPostgresQuery(filteredQuery, orderByQuery, pageSize, pageNumber, len(args))
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
