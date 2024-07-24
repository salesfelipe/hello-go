package strain

func Keep[T any](list []T, predicate func(i T) bool) []T {
	matches := 0

	for _, v := range list {
		if predicate(v) {
			matches++
		}
	}

	result := make([]T, matches)

	copyPos := 0
	for _, v := range list {
		if predicate(v) {
			result[copyPos] = v
			copyPos++
		}
	}

	return result
}

func Discard[T any](list []T, predicate func(i T) bool) []T {
	matches := 0
	for _, v := range list {
		if !predicate(v) {
			matches++
		}
	}

	result := make([]T, matches)

	copyPos := 0
	for _, v := range list {
		if !predicate(v) {
			result[copyPos] = v
			copyPos++
		}
	}

	return result
}
