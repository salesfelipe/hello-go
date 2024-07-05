package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(-|~|=|\*)*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i)password"`)

	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line+\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reContains := regexp.MustCompile(`User\s+([A-z]+\d+)`)

	result := []string{}
	for _, line := range lines {
		if reContains.MatchString(line) {
			userName := reContains.FindStringSubmatch(line)[1]

			result = append(result, fmt.Sprintf("[USR] %s ", userName)+line)

		} else {
			result = append(result, line)
		}
	}

	return result
}
