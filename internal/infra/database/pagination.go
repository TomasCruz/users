package database

import (
	"fmt"
	"strings"
)

func buildPaginatedQuery(filteredQuery, orderByQuery string, pageSize, pageNumber, argsCount int) (string, int, int) {
	if pageSize <= 0 || pageNumber <= 0 {
		return filteredQuery, 0, 0
	}

	var sb strings.Builder
	sb.WriteString(filteredQuery)
	if orderByQuery != "" {
		sb.WriteRune(' ')
		sb.WriteString(orderByQuery)
	}

	// add pagination
	limit := pageSize
	offset := limit * (pageNumber - 1)

	sb.WriteString(fmt.Sprintf(" LIMIT $%d", argsCount+1))
	sb.WriteString(fmt.Sprintf(" OFFSET $%d", argsCount+2))

	return sb.String(), limit, offset
}
