// Package: hello-package
package main

import (
	"fmt"
	"hello-go/interest"
)

// Main do stuff
func main() {

	fmt.Println(interest.YearsBeforeDesiredBalance(100, 125.8))
}
