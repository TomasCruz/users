package database

import (
	"fmt"
	"strings"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/pkg/errors"
)

func (pDB postgresDB) countFilteredQueryResults(filteredQueryString string, args []interface{}) (int64, error) {
	queryString := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS a", filteredQueryString)

	var totalCount int64
	err := pDB.db.QueryRow(queryString, args...).Scan(&totalCount)
	if err != nil {
		return 0, errors.Wrap(entities.ErrCountFilteredQuery, err.Error())
	}

	return totalCount, nil
}

func (pDB postgresDB) makeFilteredQuery(filter entities.UserFilter, basicQuery string, parameterNames map[string]string) (string, []interface{}) {
	if filter.Empty() {
		return basicQuery, nil
	}

	return pDB.addFiltersToQuery(filter, basicQuery, parameterNames)
}

func (pDB postgresDB) addFiltersToQuery(filter entities.UserFilter, basicQuery string, parameterNames map[string]string) (string, []interface{}) {
	args := []interface{}{}
	countryFilteredQuery, args := pDB.addParticularFilterToQuery(basicQuery, "country", filter.Country, parameterNames, args)
	return countryFilteredQuery, args
}

func (pDB postgresDB) addParticularFilterToQuery(inputQuery, filterName string,
	filterValues []string,
	parameterNames map[string]string,
	args []interface{}) (string, []interface{}) {

	if len(filterValues) == 0 {
		return inputQuery, args
	}

	filterParameterName, present := parameterNames[filterName]
	if !present {
		return inputQuery, args
	}

	var sb strings.Builder

	argsNextIndex := len(args) + 1
	args = append(args, filterValues[0])
	sb.WriteString(fmt.Sprintf("%s IN ($%d", filterParameterName, argsNextIndex))

	for i := 1; i < len(filterValues); i++ {
		args = append(args, filterValues[i])
		argsNextIndex++
		sb.WriteString(fmt.Sprintf(", $%d", argsNextIndex))
	}
	sb.WriteRune(')')
	currFilterPart := sb.String()

	// add filters to query
	whereIndex := strings.Index(inputQuery, "WHERE")
	sb = strings.Builder{}
	sb.WriteString(inputQuery)
	if whereIndex == -1 {
		sb.WriteString(" WHERE ")
	} else {
		sb.WriteString(" AND ")
	}
	sb.WriteString(currFilterPart)

	return sb.String(), args
}

func (pDB postgresDB) makePaginatedQuery(filteredQuery, orderBy string, paginator entities.Paginator, args []interface{}) (string, []interface{}) {
	if paginator.Empty() {
		return filteredQuery, args
	}

	var sb strings.Builder
	sb.WriteString(filteredQuery)
	if orderBy != "" {
		sb.WriteRune(' ')
		sb.WriteString(orderBy)
	}

	// add pagination
	limit := paginator.PageSize
	args = append(args, limit)

	offset := limit * (paginator.PageNumber - 1)
	args = append(args, offset)

	sb.WriteString(fmt.Sprintf(" LIMIT $%d", len(args)-1))
	sb.WriteString(fmt.Sprintf(" OFFSET $%d", len(args)))

	return sb.String(), args
}
