package entities

import (
	"net/url"
	"strconv"

	"github.com/TomasCruz/users/internal/errlog"
	"github.com/pkg/errors"
)

type Paginator struct {
	PageSize   int
	PageNumber int
}

func NewPaginator(values url.Values) (Paginator, error) {
	var pageSize int
	var pageNumber int

	// extract paginator
	for k, urlValue := range values {
		// only page size and number are valid, and there can be only one of each
		switch k {
		case "page-size":
			ps64, err := strconv.ParseInt(urlValue[0], 10, 64)
			if err != nil || ps64 <= int64(0) {
				if err != nil {
					err = errors.Wrap(ErrPageSize, err.Error())
				} else {
					err = errors.WithStack(ErrPageSize)
				}

				// log error and return, prevent invalid page size or number from reaching DB
				errlog.Error(err, "")
				return Paginator{}, err
			}

			pageSize = int(ps64)
		case "page-number":
			pn64, err := strconv.ParseInt(urlValue[0], 10, 64)
			if err != nil || pn64 <= int64(0) {
				if err != nil {
					err = errors.Wrap(ErrPageNumber, err.Error())
				} else {
					err = errors.WithStack(ErrPageNumber)
				}

				// log error and return, prevent invalid page size or number from reaching DB
				errlog.Error(err, "")
				return Paginator{}, err
			}

			pageNumber = int(pn64)
		}
	}

	return Paginator{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}, nil
}

func (p Paginator) Empty() bool {
	return p.PageSize == 0 || p.PageNumber == 0
}
