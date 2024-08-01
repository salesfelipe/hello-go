package leetmicrosoft

import (
	"strconv"
	"testing"
)

var casesNumberToWords []Case = []Case{
	{input: "1", expect: "One"},
	{input: "2", expect: "Two"},
	{input: "3", expect: "Three"},
	{input: "100", expect: "One Hundred"},
	{input: "200", expect: "Two Hundred"},
	{input: "203", expect: "Two Hundred Three"},
	{input: "2003", expect: "Two Thousand Three"},
	{input: "10023", expect: "Ten Thousand Twenty Three"},
	{input: "10123", expect: "Ten Thousand One Hundred Twenty Three"},
	{input: "100123", expect: "One Hundred Thousand One Hundred Twenty Three"},
	{input: "1234567", expect: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
}

func TestNumberTowords(t *testing.T) {

	for _, currTest := range casesNumberToWords {
		input, _ := strconv.Atoi(currTest.input)

		result := numberToWords(input)

		if result != currTest.expect {
			t.Errorf(`TestReverse = input: %s, response: %s | want: %s`, currTest.input, result, currTest.expect)
		}
	}
}
