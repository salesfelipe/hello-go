package leetmicrosoft

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	input := removeDuplicateSpaces(s)
	input = trimString(input)

	words := strings.Split(input, " ")

	result := ""

	for i := len(words) - 1; i >= 0; i-- {
		if i == len(words)-1 {
			result = words[i]
		} else {
			result += " " + words[i]
		}
	}

	return result
}

func removeDuplicateSpaces(input string) string {
	re := regexp.MustCompile(`(\s(\s)+)`)

	return re.ReplaceAllString(input, " ")
}

func trimString(input string) string {
	re := regexp.MustCompile(`(((^\s*))|(\s*$))`)

	return re.ReplaceAllString(input, "")
}
