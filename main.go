// Package: hello-package
package main

import (
	"fmt"
	"hello-go/hamming"
)

// Main do stuff
func main() {
	fmt.Println(hamming.Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}
