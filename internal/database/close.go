package database

import "github.com/pkg/errors"

func (pDB postgresDB) Close() error {
	err := pDB.db.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
