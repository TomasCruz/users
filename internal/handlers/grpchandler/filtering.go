package grpchandler

func extractParticularFilter(values []string) map[string]struct{} {
	// extract values
	valueSet := map[string]struct{}{}
	for _, currValue := range values {
		valueSet[currValue] = struct{}{}
	}

	return valueSet
}
