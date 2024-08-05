package leetmicrosoft

import (
	"reflect"
	"testing"
)

type CaseSpiral struct {
	input  [][]int
	expect []int
}

var casesSpiral []CaseSpiral = []CaseSpiral{
	{
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
}

func TestSpiral(t *testing.T) {

	for _, currTest := range casesSpiral {
		result := spiralOrder(currTest.input)

		if !reflect.DeepEqual(result, currTest.expect) {
			t.Errorf(`TestReverse = input: %v, response: %v | want: %v`, currTest.input, result, currTest.expect)
		}
	}
}
