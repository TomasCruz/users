package database

import (
	"fmt"
	"strings"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/pkg/errors"
)

func (pDB postgresDB) countFilteredQueryResults(filteredQueryString string, args []interface{}) (int64, error) {
	var sb strings.Builder
	sb.WriteString("SELECT COUNT(*) FROM (")
	sb.WriteString(filteredQueryString)
	sb.WriteString(") AS a")
	queryString := sb.String()

	var totalCount int64
	err := pDB.db.QueryRow(queryString, args...).Scan(&totalCount)
	if err != nil {
		return 0, errors.Wrap(entities.ErrCountFilteredQuery, err.Error())
	}

	return totalCount, nil
}

func (pDB postgresDB) makeFilteredQuery(basicQuery string, filter entities.Filter) (string, []interface{}) {
	if filter.Empty() {
		return basicQuery, nil
	}

	// make filterPart
	return pDB.addFilterToQuery(basicQuery, "country", filter.Country)
}

func (pDB postgresDB) addFilterToQuery(inputString, filterName string, filterValues []string) (string, []interface{}) {
	if len(filterValues) == 0 {
		return inputString, nil
	}

	var args []interface{}
	var sb strings.Builder

	args = append(args, filterValues[0])
	sb.WriteString(fmt.Sprintf("%s IN ($1", filterName))

	for i := 1; i < len(filterValues); i++ {
		args = append(args, filterValues[i])
		sb.WriteString(fmt.Sprintf(", $%d", i+1))
	}
	sb.WriteRune(')')
	currFilterPart := sb.String()

	// add filters to query
	whereIndex := strings.Index(inputString, "WHERE")
	sb = strings.Builder{}
	sb.WriteString(inputString)
	if whereIndex == -1 {
		sb.WriteString(" WHERE ")
	} else {
		sb.WriteString(" AND ")
	}
	sb.WriteString(currFilterPart)

	return sb.String(), args
}

func (pDB postgresDB) makePaginatedQuery(filteredQuery string, paginator entities.Paginator, args []interface{}) (string, []interface{}) {
	if paginator.Empty() {
		return filteredQuery, args
	}

	// add pagination
	limit := paginator.PageSize
	args = append(args, limit)

	offset := limit * (paginator.PageNumber - 1)
	args = append(args, offset)

	var sb strings.Builder
	sb.WriteString(filteredQuery)
	sb.WriteString(fmt.Sprintf(" LIMIT $%d", len(args)-1))
	sb.WriteString(fmt.Sprintf(" OFFSET $%d", len(args)))

	return sb.String(), args
}
