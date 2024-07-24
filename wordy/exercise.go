package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

const add string = "plus"
const subtract string = "minus"
const multiply string = "multiplied by"
const divide string = "divided by"

func Answer(question string) (int, bool) {
	err := validateQuestion(question)
	if err != nil {
		return 0, false
	}

	numbers := extractNumbers(question)
	operations := extractOperations(question)

	result := numbers[0]

	for i := 0; i < len(operations); i++ {
		result = executeOperation(operations[i], result, numbers[i+1])
	}

	if len(numbers) == 1 {
		return numbers[0], true
	}

	return result, true
}

func validateQuestion(question string) error {
	re := regexp.MustCompile(`^What is -*\d+( (plus|minus|multiplied by|divided by)+ -*\d+)*\?$`)

	if !re.MatchString(question) {
		return fmt.Errorf("invalid question")

	}

	return nil
}

func executeOperation(operation string, leftOperand int, rightOperand int) int {
	if operation == add {
		return leftOperand + rightOperand
	}

	if operation == subtract {
		return leftOperand - rightOperand
	}

	if operation == multiply {
		return leftOperand * rightOperand
	}

	if operation == divide {
		return leftOperand / rightOperand
	}

	panic("invalid operation")
}

func extractNumbers(question string) []int {
	re := regexp.MustCompile("(-*[0-9]+)")

	hits := re.FindAllString(question, -1)

	if hits == nil {
		return nil
	}

	result := make([]int, len(hits))

	for i, v := range hits {
		convertedN, err := strconv.Atoi(v)

		if err != nil {
			// This shouldn't happen because of the regex
			panic(err)
		}

		result[i] = convertedN
	}

	return result
}

func extractOperations(question string) []string {
	reg := fmt.Sprintf("(%s|%s|%s|%s)+", add, subtract, multiply, divide)
	re := regexp.MustCompile(reg)

	hits := re.FindAllString(question, -1)

	if hits == nil {
		return nil
	}

	return hits
}
