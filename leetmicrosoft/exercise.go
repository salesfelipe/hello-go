package leetmicrosoft

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func dequeue(queue []*Node) (*Node, []*Node) {
	item := queue[0]
	if len(queue) == 1 {
		return item, []*Node{}
	}

	return item, queue[1:]
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	copiedNodes := make(map[int]*Node)

	// Shallow  copy
	queue := make([]*Node, 0)
	queue = append(queue, head)
	for len(queue) > 0 {
		currentNode, q := dequeue(queue)

		queue = q

		_, ok := copiedNodes[currentNode.Val]

		if !ok {
			copy := Node{Val: currentNode.Val}
			copiedNodes[currentNode.Val] = &copy

			if currentNode.Next != nil {
				queue = append(queue, currentNode.Next)
			}
		}
	}

	// Deepy copy

	deepyCopyQ := make([]*Node, 0)
	deepyCopyQ = append(deepyCopyQ, head)
	for len(deepyCopyQ) > 0 {
		currentNode, q := dequeue(deepyCopyQ)
		deepyCopyQ = q

		currentNodeCopy, currCopyOk := copiedNodes[currentNode.Val]

		if currentNode.Next != nil {
			nextNodeCopy, nextCopyOk := copiedNodes[currentNode.Next.Val]

			if currCopyOk && nextCopyOk {
				currentNodeCopy.Next = nextNodeCopy
			}

			q = append(deepyCopyQ, currentNode.Next)

			deepyCopyQ = q
		}

		if currentNode.Random != nil {
			randomNodeCopy, randomCopyOk := copiedNodes[currentNode.Random.Val]

			if currCopyOk && randomCopyOk {
				currentNodeCopy.Random = randomNodeCopy
			}
		}
	}

	return copiedNodes[head.Val]
}

func TraverseNPrint(head *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, head)
	for len(queue) > 0 {
		currentNode, q := dequeue(queue)
		queue = q

		printNode(currentNode)

		if currentNode.Next != nil {
			queue = append(queue, currentNode.Next)
		}
	}

}

func printNode(node *Node) {

	next := -1
	random := -1

	if node.Next != nil {
		next = node.Next.Val
	}

	if node.Random != nil {
		random = node.Random.Val
	}

	fmt.Println(node.Val, next, random)
}
