// Package: hello-package
package main

import (
	"fmt"
	"hello-go/gigasecond"
	"time"
)

// Main do stuff
func main() {
	fmt.Println(gigasecond.AddGigasecond(time.Date(2015, 1, 24, 22, 0, 0, 0, time.UTC)))
}
