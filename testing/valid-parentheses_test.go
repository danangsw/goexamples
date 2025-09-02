package testing_test

import (
	"fmt"
	"testing"

	problem "../problem"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"[", false},
		{"", false},
		{"({[([{}])]})", true},
		{"]({[([{}])]})", false},
		{"][", false},
		{"[][]]", false},
		{"{[[]}", false},
		{"{}{}{}{}[[[()]]]", true},
	}

	fmt.Println("Testing IsValidParentheses:")
	for _, test := range tests {
		t.Run(fmt.Sprintf("IsValidParentheses(%q):\n", test.input), func(t *testing.T) {
			result := problem.IsValidParentheses(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
