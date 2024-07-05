package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.

// recommendation, search, weather, default
func Application(log string) string {
	applicationsMap := make(map[rune]string)

	applicationsMap['❗'] = "recommendation"
	applicationsMap['🔍'] = "search"
	applicationsMap['☀'] = "weather"

	for _, char := range log {
		application, exists := applicationsMap[char]

		if exists {
			return application
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""

	for _, char := range log {
		if char == oldRune {
			result += string(newRune)
		} else {
			result += string(char)
		}
	}

	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
