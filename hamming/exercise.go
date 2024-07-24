package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("streams with different sizes")
	}

	result := 0
	for pos := range a {
		if a[pos] != b[pos] {
			result += 1
		}

	}

	return result, nil
}
