package leetmicrosoft

import (
	"reflect"
	"testing"
)

type CaseAddTw struct {
	inputL1 []int
	inputL2 []int
	expect  []int
}

var casesAddTw []CaseAddTw = []CaseAddTw{
	{inputL1: []int{2, 4, 3}, inputL2: []int{5, 6, 4}, expect: []int{7, 0, 8}},
	{inputL1: []int{0}, inputL2: []int{0}, expect: []int{0}},
	{inputL1: []int{9, 9, 9, 9, 9, 9, 9}, inputL2: []int{9, 9, 9, 9}, expect: []int{8, 9, 9, 9, 0, 0, 0, 1}},
	{inputL1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, inputL2: []int{5, 6, 4}, expect: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
}

func TestAddTwoSum(t *testing.T) {

	for _, currTest := range casesAddTw {
		l1 := convertToLinkedList(currTest.inputL1)
		l2 := convertToLinkedList(currTest.inputL2)

		result := addTwoNumbers(l1, l2)

		resultArray := extractListValues(result)

		if !reflect.DeepEqual(resultArray, currTest.expect) {
			t.Errorf(`TestAddTwoSum = input: %v, response: %v | want: %v`, currTest.inputL1, resultArray, currTest.expect)
		}
	}
}
