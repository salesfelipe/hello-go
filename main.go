// Package: hello-package
package main

import (
	"hello-go/leetmicrosoft"
)

// Main do stuff
func main() {
	first := leetmicrosoft.Node{
		Val: 1,
	}

	second := leetmicrosoft.Node{
		Val: 2,
	}

	third := leetmicrosoft.Node{
		Val: 3,
	}

	first.Next = &second
	first.Random = &third
	second.Next = &third

	copyH := leetmicrosoft.CopyRandomList(&first)

	println(copyH.Val, first.Val, copyH.Next.Val, first.Next.Val)
}
