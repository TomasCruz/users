package database

import (
	"database/sql"
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (pDB postgresDB) GetUserByID(userID uuid.UUID) (entities.User, error) {
	var (
		firstName string
		lastName  string
		pswdHash  string
		email     string
		country   string
		createdAt time.Time
		updatedAt time.Time
	)

	queryString := `SELECT first_name, last_name, pswd_hash, email, country, created_at, updated_at
		FROM users
		WHERE user_id=$1`

	err := pDB.db.QueryRow(queryString, userID).
		Scan(&firstName, &lastName, &pswdHash, &email, &country, &createdAt, &updatedAt)
	if err != nil {
		// sql.ErrNoRows -> ErrNonexistingUser
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.WithStack(entities.ErrNonexistingUser)
		}

		return entities.User{}, errors.Wrap(entities.ErrGetUser, err.Error())
	}

	return entities.User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		PswdHash:  pswdHash,
		Email:     email,
		Country:   country,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (pDB postgresDB) GetUserByEmail(email string) (entities.User, error) {
	var (
		userID    uuid.UUID
		firstName string
		lastName  string
		pswdHash  string
		country   string
		createdAt time.Time
		updatedAt time.Time
	)

	queryString := `SELECT user_id, first_name, last_name, pswd_hash, country, created_at, updated_at
		FROM users
		WHERE email=$1`

	err := pDB.db.QueryRow(queryString, email).
		Scan(&userID, &firstName, &lastName, &pswdHash, &country, &createdAt, &updatedAt)
	if err != nil {
		// sql.ErrNoRows -> ErrNonexistingUser
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.WithStack(entities.ErrNonexistingUser)
		}

		return entities.User{}, errors.Wrap(entities.ErrGetUser, err.Error())
	}

	return entities.User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		PswdHash:  pswdHash,
		Email:     email,
		Country:   country,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (pDB postgresDB) getUserByIDForUpdate(userID uuid.UUID) (entities.User, error) {
	var (
		firstName string
		lastName  string
		pswdHash  string
		email     string
		country   string
		createdAt time.Time
		updatedAt time.Time
	)

	queryString := `SELECT first_name, last_name, pswd_hash, email, country, created_at, updated_at
		FROM users
		WHERE user_id=$1
		FOR NO KEY UPDATE`

	err := pDB.db.QueryRow(queryString, userID).
		Scan(&firstName, &lastName, &pswdHash, &email, &country, &createdAt, &updatedAt)
	if err != nil {
		// sql.ErrNoRows -> ErrNonexistingUser
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, errors.WithStack(entities.ErrNonexistingUser)
		}

		return entities.User{}, errors.Wrap(entities.ErrGetUser, err.Error())
	}

	return entities.User{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		PswdHash:  pswdHash,
		Email:     email,
		Country:   country,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
