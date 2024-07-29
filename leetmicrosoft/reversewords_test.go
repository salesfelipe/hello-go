package leetmicrosoft

import (
	"testing"
)

type Case struct {
	input  string
	expect string
}

var cases []Case = []Case{
	{input: "the    sky is blue", expect: "blue is sky the"},
	{input: "  hello world  ", expect: "world hello"},
	{input: "a good   example", expect: "example good a"},
}

func TestReverse(t *testing.T) {

	for _, currTest := range cases {
		result := reverseWords(currTest.input)

		if result != currTest.expect {
			t.Errorf(`TestReverse = input: %s, response: %s | want: %s`, currTest.input, result, currTest.expect)
		}
	}
}
