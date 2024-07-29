package leetmicrosoft

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func getKey(node *Node) string {
	return fmt.Sprintf("%p", node)
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
	copiedNodes := make(map[string]*Node)

	// Shallow  copy
	queue := make([]*Node, 0)
	queue = append(queue, head)

	// a := fmt.Stringer(head.Next)
	for len(queue) > 0 {
		currentNode, q := dequeue(queue)
		nodeKey := getKey(currentNode)

		queue = q

		_, ok := copiedNodes[nodeKey]

		if !ok {
			copy := Node{Val: currentNode.Val}
			copiedNodes[nodeKey] = &copy

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
		currentKey := getKey(currentNode)

		currentNodeCopy, currCopyOk := copiedNodes[currentKey]

		if currentNode.Next != nil {
			nextKey := getKey(currentNode.Next)
			nextNodeCopy, nextCopyOk := copiedNodes[nextKey]

			if currCopyOk && nextCopyOk {
				currentNodeCopy.Next = nextNodeCopy
			}

			q = append(deepyCopyQ, currentNode.Next)

			deepyCopyQ = q
		}

		if currentNode.Random != nil {
			randomKey := getKey(currentNode.Random)
			randomNodeCopy, randomCopyOk := copiedNodes[randomKey]

			if currCopyOk && randomCopyOk {
				currentNodeCopy.Random = randomNodeCopy
			}
		}
	}

	return copiedNodes[getKey(head)]
}

func TraverseNPrint(head *Node) {

	var currentNode *Node = head

	for currentNode.Next != nil {
		printNode(currentNode)
		printNode(currentNode.Next)

		currentNode = currentNode.Next
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
