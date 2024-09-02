package database

import (
	"database/sql"

	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// file driver for migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// InitDB does DB migrations and verifies DB accessibility
func InitDB(config configuration.Config) (core.DB, error) {
	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "", driver)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			err = nil
		} else {
			return nil, errors.WithStack(err)
		}
	}

	sErr, dbErr := m.Close()
	if sErr != nil {
		return nil, errors.WithStack(sErr)
	} else if dbErr != nil {
		return nil, errors.WithStack(dbErr)
	}

	db, err = openAndCheck(config.DbURL)
	if err != nil {
		return nil, err
	}

	return postgresDB{db: db}, nil
}

func openAndCheck(dbString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

type postgresDB struct {
	db *sql.DB
}

// type postgresTx struct {
// 	*sql.Tx
// }
