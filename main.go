// Package: hello-package
package main

import (
	"fmt"
	"hello-go/parsinglogfiles"
)

// Main do stuff
func main() {
	fmt.Println(parsinglogfiles.CountQuotedPasswords([]string{
		``,
		`[INF] passWord`,
		`"passWord"`,
		`[INF] User saw error message "Unexpected Error" on page load.`,
		`[INF] The message "Please reset your password" was ignored by the user`,
	}))
}
