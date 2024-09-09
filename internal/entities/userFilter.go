package entities

func ExtractUserFilter(filter map[string][]string) map[string][]string {
	// extract filter
	var countries []string
	for k, v := range filter {
		if !validUserFilterKey(k) {
			continue
		}

		switch k {
		case "country":
			for _, v := range v {
				if len(v) != 2 && len(v) != 3 {
					// continue, invalid filter simply won't be applied
					continue
				}
				countries = append(countries, v)
			}
		}
	}

	userFilter := map[string][]string{}
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
