package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("not a valid input")
	}

	if n == 1 {
		return 0, nil
	}

	isEven := n%2 == 0

	nextCycle := 0
	if isEven {
		nextCycle = n / 2
	} else {
		nextCycle = n*3 + 1
	}

	result, err := CollatzConjecture(nextCycle)

	return (1 + result), err
}
