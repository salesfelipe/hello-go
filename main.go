// Package: hello-package
package main

import (
	"fmt"
	"hello-go/leetmicrosoft"
)

// Main do stuff
func main() {
	// [[7,null],[13,0],[11,4],[10,2],[1,0]]

	cache := leetmicrosoft.Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

}
