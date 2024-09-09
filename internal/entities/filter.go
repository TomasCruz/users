package entities

import (
	"fmt"
	"strings"
)

func ExtractFilter(values map[string][]string) map[string][]string {
	filter := map[string][]string{}

	// extract filter
	for k, urlValue := range values {
		valueSet := map[string]struct{}{}
		for _, uv := range urlValue {
			if strings.Contains(uv, ",") {
				currValues := strings.Split(uv, ",")
				for _, vs := range currValues {
					valueSet[vs] = struct{}{}
				}
			} else {
				valueSet[uv] = struct{}{}
			}
		}

		values := []string{}
		for v := range valueSet {
			values = append(values, v)
		}

		filter[k] = values
	}

	return filter
}

func BuildFilteredPostgresQuery(basicQuery string, filter map[string][]string, parameterNames map[string]string) (string, []interface{}) {
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

func applyParticularFilterToPostgresQuery(inputQuery, filterParameterName string, filterValues []string, args []interface{}) (string, []interface{}) {
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
