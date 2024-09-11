package database

import (
	"fmt"
	"strings"
)

func buildFilteredQuery(basicQuery string, filter map[string]map[string]struct{}, parameterNames map[string]string) (string, []interface{}) {
	if len(filter) == 0 {
		return basicQuery, nil
	}

	args := []interface{}{}
	filteredQuery := basicQuery

	for k, v := range filter {
		filterParameterName, present := parameterNames[k]
		if present {
			filteredQuery, args = applyParticularFilterToPostgresQuery(filteredQuery, filterParameterName, v, args)
		}
	}

	return filteredQuery, args
}

func applyParticularFilterToPostgresQuery(inputQuery, filterParameterName string, filterValues map[string]struct{}, args []interface{}) (string, []interface{}) {
	var sb strings.Builder

	first := true
	argsNextIndex := len(args) + 1
	for fv := range filterValues {
		if first {
			args = append(args, fv)
			sb.WriteString(fmt.Sprintf("%s IN ($%d", filterParameterName, argsNextIndex))
			first = false
			continue
		}
		args = append(args, fv)
		argsNextIndex++
		sb.WriteString(fmt.Sprintf(", $%d", argsNextIndex))
	}
	sb.WriteRune(')')
	currFilterPart := sb.String()

	// add filter to query
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
