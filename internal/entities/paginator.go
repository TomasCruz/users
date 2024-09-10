package entities

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractPagination(values map[string][]string, defaultPageSize, defaultPageNumber *int) (int, int) {
	var pageSize int
	var pageNumber int

	// extract pagination
	for k, urlValue := range values {
		switch k {
		case "page-size":
			// consider only one page size value
			ps64, err := strconv.ParseInt(urlValue[0], 10, 64)
			if err == nil {
				pageSize = int(ps64)
			}
		case "page-number":
			// consider only one page number value
			pn64, err := strconv.ParseInt(urlValue[0], 10, 64)
			if err == nil {
				pageNumber = int(pn64)
			}
		}
	}

	// for invalid pagination values
	if pageSize <= 0 || pageNumber <= 0 {
		if defaultPageSize != nil && defaultPageNumber != nil {
			// use default
			return *defaultPageSize, *defaultPageNumber
		} else {
			// don't apply pagination
			return 0, 0
		}
	}

	return pageSize, pageNumber
}

func BuildPaginatedPostgresQuery(filteredQuery, orderByQuery string, pageSize, pageNumber, argsCount int) (string, int, int) {
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
