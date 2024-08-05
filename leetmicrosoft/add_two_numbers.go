package leetmicrosoft

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Number := getNumberFromList(l1)
	l2Number := getNumberFromList(l2)

	sum := l1Number + l2Number

	return generateListFromNumber(sum)
}

func getNumberFromList(list *ListNode) int {
	node := list
	extractedNumber := ""
	for node != nil {
		extractedNumber = fmt.Sprintf("%d%s", node.Val, extractedNumber)
		node = node.Next
	}

	value, err := strconv.Atoi(extractedNumber)

	if err != nil {
		return 0
	}

	return value
}

func generateListFromNumber(number int) *ListNode {

	stringifiedVersion := fmt.Sprintf("%d", number)
	digits := strings.Split(stringifiedVersion, "")

	var currentNode *ListNode
	var head *ListNode
	for i := len(digits) - 1; i >= 0; i-- {
		value, _ := strconv.Atoi(digits[i])
		newNode := &ListNode{Val: value}

		if currentNode != nil {
			currentNode.Next = newNode
		} else {
			head = newNode
		}

		currentNode = newNode
	}

	return head
}

func extractListValues(list *ListNode) []int {
	digits := make([]int, 0)
	node := list
	for node != nil {
		digits = append(digits, node.Val)
		node = node.Next
	}

	return digits
}

func convertToLinkedList(digits []int) *ListNode {
	var currentNode *ListNode
	var head *ListNode
	for i := 0; i < len(digits); i++ {
		newNode := &ListNode{Val: digits[i]}

		if currentNode != nil {
			currentNode.Next = newNode
		} else {
			head = newNode
		}

		currentNode = newNode
	}

	return head
}
