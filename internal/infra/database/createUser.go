package database

import (
	"time"

	"github.com/TomasCruz/users/internal/domain/entities"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (pDB postgresDB) CreateUser(userID uuid.UUID, firstName, lastName, pswdHash, email, country string, createdAt, updatedAt time.Time) (entities.User, error) {
	pTx, err := pDB.newTransaction()
	if err != nil {
		return entities.User{}, err
	}
	defer pTx.commitOrRollbackOnError(&err)

	if _, err := pDB.getUserByEmail(email); err != nil {
		// ignore no rows error
		if !errors.Is(err, entities.ErrNonexistingUser) {
			return entities.User{}, err
		}
	} else {
		// no error, user found by email
		return entities.User{}, entities.ErrExistingEmail
	}

	sqlStatement := `INSERT INTO users (user_id, first_name, last_name, pswd_hash, email, country, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := pTx.Prepare(sqlStatement)
	if err != nil {
		return entities.User{}, errors.WithStack(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, firstName, lastName, pswdHash, email, country, createdAt, updatedAt); err != nil {
		return entities.User{}, errors.WithStack(err)
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
