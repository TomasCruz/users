package entities

import (
	"net/url"
	"strings"

	"github.com/TomasCruz/users/internal/errlog"
)

type UserFilter struct {
	Country []string
}

func New(values url.Values) UserFilter {
	filter := UserFilter{}

	// extract filter
	for k, urlValue := range values {
		if !filter.ValidFilterKey(k) {
			continue
		}

		var vs []string
		for _, uv := range urlValue {
			if strings.Contains(uv, ",") {
				currVs := strings.Split(uv, ",")
				vs = append(vs, currVs...)
			} else {
				vs = append(vs, uv)
			}
		}

		switch k {
		case "country":
			for _, v := range vs {
				if len(v) != 2 && len(v) != 3 {
					// log error and continue, invalid filter simply won't be applied
					errlog.Error(ErrCountryLength, "")
					continue
				}
				filter.Country = append(filter.Country, v)
			}
		}
	}

	return filter
}

func (f UserFilter) Empty() bool {
	return len(f.Country) == 0
}

func (f UserFilter) ValidFilterKey(key string) bool {
	switch key {
	case "country":
		return true
	default:
		return false
	}
}
