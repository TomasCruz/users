package httphandler

import (
	"net/url"
	"strings"
)

func extractFilterAndPaginatorFromQueryParams(values url.Values) map[string]map[string]struct{} {
	filter := map[string]map[string]struct{}{}

	// extract filter
	for k, urlValue := range values {
		valueSet := map[string]struct{}{}
		for _, uv := range urlValue {
			if strings.Contains(uv, ",") {
				currValues := strings.Split(uv, ",")
				for _, vs := range currValues {
					if vs != "" {
						valueSet[vs] = struct{}{}
					}
				}
			} else {
				if uv != "" {
					valueSet[uv] = struct{}{}
				}
			}
		}

		if len(valueSet) != 0 {
			filter[k] = valueSet
		}
	}

	return filter
}
