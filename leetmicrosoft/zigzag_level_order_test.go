package leetmicrosoft

import (
	"reflect"
	"testing"
)

type CaseZigZag struct {
	inputL1 []int
	expect  [][]int
}

var casesZigZag []CaseZigZag = []CaseZigZag{
	{inputL1: []int{2, 4, 0, 3}, expect: [][]int{{7, 2}, {7, 2}, {7, 3}}},
}

func TestZigZag(t *testing.T) {

	for _, currTest := range casesZigZag {
		l1 := convertToLinkedList(currTest.inputL1)

		result := addTwoNumbers(l1, l2)

		resultArray := extractListValues(result)

		if !reflect.DeepEqual(resultArray, currTest.expect) {
			t.Errorf(`TestAddTwoSum = input: %v, response: %v | want: %v`, currTest.inputL1, resultArray, currTest.expect)
		}
	}
}
