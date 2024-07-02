// Package: hello-package
package main

import (
	"fmt"
	"hello-go/electionday"
)

// Main do stuff
func main() {
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	electionday.DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])

}
