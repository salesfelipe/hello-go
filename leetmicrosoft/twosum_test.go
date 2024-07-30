package leetmicrosoft

import (
	"testing"
)

type CaseTw struct {
	inputNums   []int
	inputTarget int
	expect      []int
}

var casesTw []CaseTw = []CaseTw{
	{inputNums: []int{2, 7, 11, 15}, inputTarget: 9, expect: []int{0, 1}},
	{inputNums: []int{3, 2, 4}, inputTarget: 6, expect: []int{1, 2}},
	{inputNums: []int{3, 3}, inputTarget: 6, expect: []int{0, 1}},
}

func TestTwoSum(t *testing.T) {

	for _, currTest := range casesTw {
		result := twoSum(currTest.inputNums, currTest.inputTarget)

		if result[0] != currTest.expect[0] || result[1] != currTest.expect[1] {
			t.Errorf(`TestReverse = input: %v, response: %v | want: %v`, currTest.inputNums, result, currTest.expect)
		}
	}
}
