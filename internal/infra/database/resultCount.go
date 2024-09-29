package database

import (
	"fmt"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/pkg/errors"
)

func (pDB postgresDB) resultCountQuery(filteredQuery string, args []interface{}) (int64, error) {
	countFilteredQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS a", filteredQuery)

	var totalCount int64
	err := pDB.db.QueryRow(countFilteredQuery, args...).Scan(&totalCount)
	if err != nil {
		return 0, errors.Wrap(entities.ErrCountFilteredQuery, err.Error())
	}

	return totalCount, nil
}
