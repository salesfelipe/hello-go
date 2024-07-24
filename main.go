// Package: hello-package
package main

import (
	"fmt"
	"hello-go/strain"
)

// Main do stuff
func main() {
	has5 := func(list []int) bool {
		for _, entry := range list {
			if entry == 5 {
				return true
			}
		}
		return false
	}

	list := [][]int{
		{1, 2, 3},
		{5, 5, 5},
		{5, 1, 2},
		{2, 1, 2},
		{1, 5, 2},
		{2, 2, 1},
		{1, 2, 5},
	}

	fmt.Println(strain.Keep(list, has5))
}
