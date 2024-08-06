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

	l1Number := extractListValues(l1)
	l2Number := extractListValues(l2)

	maxLen := len(l1Number)

	if len(l2Number) > maxLen {
		maxLen = len(l2Number)
	}

	sumDigits := make([]int, 0)
	rest := 0
	for i := 0; i < maxLen; i++ {
		l1Digit := 0
		l2Digit := 0

		if i < len(l1Number) {
			l1Digit = l1Number[i]
		}

		if i < len(l2Number) {
			l2Digit = l2Number[i]
		}

		sum := l1Digit + l2Digit + rest

		if rest != 0 {
			rest = 0
		}

		if sum > 9 {
			rest = 1
			sumDigits = append(sumDigits, sum-10)
		} else {
			sumDigits = append(sumDigits, sum)
		}
	}

	if rest > 0 {
		sumDigits = append(sumDigits, rest)
	}

	return convertToLinkedList(sumDigits)
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
