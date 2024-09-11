package adapters

func ExtractUserFilter(filter map[string]map[string]struct{}) map[string]map[string]struct{} {
	// extract filter
	countries := map[string]struct{}{}
	for k, v := range filter {
		if !validUserFilterKey(k) {
			continue
		}

		switch k {
		case "country":
			for v := range v {
				if len(v) != 2 && len(v) != 3 {
					// continue, invalid filter simply won't be applied
					continue
				}
				countries[v] = struct{}{}
			}
		}
	}

	userFilter := map[string]map[string]struct{}{}
	if len(countries) != 0 {
		userFilter["country"] = countries
	}

	return userFilter
}

func validUserFilterKey(key string) bool {
	switch key {
	case "country":
		return true
	default:
		return false
	}
}
