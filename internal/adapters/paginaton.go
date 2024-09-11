package adapters

import (
	"strconv"
)

func ExtractPagination(values map[string]map[string]struct{}, defaultPageSize, defaultPageNumber *int) (int, int) {
	var pageSize int
	var pageNumber int

	// extract pagination
	for k, urlValue := range values {
		switch k {
		case "page-size":
			// consider only one page size value
			for u := range urlValue {
				p64, err := strconv.ParseInt(u, 10, 64)
				if err == nil {
					pageSize = int(p64)
				}
				break
			}
		case "page-number":
			// consider only one page number value
			for u := range urlValue {
				p64, err := strconv.ParseInt(u, 10, 64)
				if err == nil {
					pageNumber = int(p64)
				}
				break
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
