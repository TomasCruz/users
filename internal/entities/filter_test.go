//go:build unit
// +build unit

package entities

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFilterFromQueryParams(t *testing.T) {
	urlStart := "http://localhost:1234/users"
	tests := []struct {
		name      string
		urlString string
		urlValues url.Values
		m         map[string][]string
	}{
		{
			name:      "no filter",
			urlString: urlStart,
			urlValues: url.Values{},
			m:         map[string][]string{},
		},
		{
			name:      "country only",
			urlString: fmt.Sprintf("%s?country=BIH", urlStart),
			urlValues: url.Values{"country": []string{"BIH"}},
			m:         url.Values{"country": []string{"BIH"}},
		},
		{
			name:      "invalid country",
			urlString: fmt.Sprintf("%s?country=", urlStart),
			urlValues: url.Values{"country": []string{""}},
			m:         map[string][]string{},
		},
		{
			name:      "country duplicated",
			urlString: fmt.Sprintf("%s?country=BIH,BIH&country=BIH", urlStart),
			urlValues: url.Values{"country": []string{"BIH,BIH", "BIH"}},
			m:         map[string][]string{"country": {"BIH"}},
		},
		{
			name:      "pagination only",
			urlString: fmt.Sprintf("%s?page-number=1&page-size=10", urlStart),
			urlValues: url.Values{"page-number": []string{"1"}, "page-size": []string{"10"}},
			m:         map[string][]string{"page-number": {"1"}, "page-size": {"10"}},
		},
		{
			name:      "pagination duplicated",
			urlString: fmt.Sprintf("%s?page-number=1,1&page-size=10&page-number=1", urlStart),
			urlValues: url.Values{"page-number": []string{"1,1", "1"}, "page-size": []string{"10"}},
			m:         map[string][]string{"page-number": {"1"}, "page-size": {"10"}},
		},
		{
			name:      "pagination bad",
			urlString: fmt.Sprintf("%s?page-number=&page-size=1", urlStart),
			urlValues: url.Values{"page-number": {""}, "page-size": {"1"}},
			m:         map[string][]string{"page-size": {"1"}},
		},
		{
			name:      "pagination and country",
			urlString: fmt.Sprintf("%s?page-number=1&page-size=10&country=SRB", urlStart),
			urlValues: url.Values{"page-number": {"1"}, "page-size": {"10"}, "country": {"SRB"}},
			m:         map[string][]string{"page-number": {"1"}, "page-size": {"10"}, "country": {"SRB"}},
		},
		{
			name:      "pagination and country lot of it",
			urlString: fmt.Sprintf("%s?page-number=1&&country=BIH&page-size=10&page-number=12&country=SRB,HRV", urlStart),
			urlValues: url.Values{"page-number": {"1", "12"}, "page-size": {"10"}, "country": {"BIH", "SRB,HRV"}},
			m:         map[string][]string{"page-number": {"1", "12"}, "page-size": {"10"}, "country": {"BIH", "SRB", "HRV"}},
		},
	}

	for _, tt := range tests {
		u, _ := url.Parse(tt.urlString)
		urlValues := u.Query()
		assert.Condition(t, func() bool { return equalMaps(tt.urlValues, urlValues) }, tt.name, "equal query param maps expected")

		m := ExtractFilterFromQueryParams(urlValues)
		assert.Condition(t, func() bool { return equalMaps(tt.m, m) }, tt.name, "equal filtered query param maps expected")
	}
}

func equalMaps(expected, actual map[string][]string) bool {
	if expected == nil {
		return actual == nil
	}

	if actual == nil {
		return false
	}

	lActual := len(actual)
	if len(expected) != lActual {
		return false
	}

	for k, v := range actual {
		vExpected, present := expected[k]
		if !present {
			return false
		}

		l := len(v)
		if l != len(vExpected) {
			return false
		}

		vSet := map[string]struct{}{}
		for i := 0; i < l; i++ {
			vSet[v[i]] = struct{}{}
		}

		vExpSet := map[string]struct{}{}
		for i := 0; i < l; i++ {
			vExpSet[vExpected[i]] = struct{}{}
		}

		for k := range vSet {
			_, present := vExpSet[k]
			if !present {
				return false
			}
		}
	}

	return true
}
