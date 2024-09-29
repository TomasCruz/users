package database

import (
	"time"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (pDB postgresDB) CreateUser(req entities.UserDTO, userID uuid.UUID, createdAt, updatedAt time.Time) (entities.User, error) {
	pTx, err := pDB.newTransaction()
	if err != nil {
		return entities.User{}, err
	}
	defer pTx.commitOrRollbackOnError(&err)

	sqlStatement := `INSERT INTO users (user_id, first_name, last_name, pswd_hash, email, country, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := pTx.Prepare(sqlStatement)
	if err != nil {
		return entities.User{}, errors.WithStack(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, *req.FirstName, *req.LastName, *req.PswdHash, *req.Email, *req.Country, createdAt, updatedAt); err != nil {
		return entities.User{}, errors.WithStack(err)
	}

	return entities.User{
		UserID:    userID,
		FirstName: *req.FirstName,
		LastName:  *req.LastName,
		PswdHash:  *req.PswdHash,
		Email:     *req.Email,
		Country:   *req.Country,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
