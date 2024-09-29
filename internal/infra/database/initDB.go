package database

import (
	"database/sql"

	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// file driver for migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// InitDB does DB migrations and verifies DB accessibility
func InitDB(config configuration.Config, logger ports.Logger) (ports.DB, error) {
	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m, err := migrate.NewWithDatabaseInstance(config.DBMigrationPath, "", driver)
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

	db, err = openAndCheck(config.DBURL)
	if err != nil {
		return nil, err
	}

	return postgresDB{
		db:     db,
		config: config,
		logger: logger,
	}, nil
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
	db     *sql.DB
	config configuration.Config
	logger ports.Logger
}

type postgresTx struct {
	*sql.Tx
}
