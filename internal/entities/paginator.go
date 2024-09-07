package entities

// type Paginator struct {
// 	PageSize   int
// 	PageNumber int
// }

// func ExtractPagination(values url.Values, withDefault *Paginator) Paginator {
// 	var pageSize int
// 	var pageNumber int

// 	// extract paginator
// 	for k, urlValue := range values {
// 		switch k {
// 		case "page-size":
// 			// consider only one page size value
// 			ps64, err := strconv.ParseInt(urlValue[0], 10, 64)
// 			if err == nil {
// 				pageSize = int(ps64)
// 			}
// 		case "page-number":
// 			// consider only one page number value
// 			pn64, err := strconv.ParseInt(urlValue[0], 10, 64)
// 			if err == nil {
// 				pageNumber = int(pn64)
// 			}
// 		}
// 	}

// 	p := Paginator{
// 		PageSize:   pageSize,
// 		PageNumber: pageNumber,
// 	}

// 	// for invalid pagination values
// 	if !p.valid() {
// 		if withDefault != nil {
// 			// use default
// 			return *withDefault
// 		} else {
// 			// don't apply pagination
// 			return Paginator{}
// 		}
// 	}

// 	return p
// }

// func (p Paginator) valid() bool {
// 	return p.PageSize > 0 && p.PageNumber > 0
// }
