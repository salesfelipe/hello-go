// Package: hello-package
package main

import (
	"hello-go/leetmicrosoft"
)

// Main do stuff
func main() {
	// [[7,null],[13,0],[11,4],[10,2],[1,0]]
	first := leetmicrosoft.Node{
		Val: 3,
	}

	second := leetmicrosoft.Node{
		Val: 3,
	}

	first.Next = &second

	copyH := leetmicrosoft.CopyRandomList(&first)

	leetmicrosoft.TraverseNPrint(copyH)
}
