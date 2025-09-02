package testing_test

import (
	"fmt"
	"testing"

	problem "../problem"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"", "b"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "", "aa"}, ""},
	}

	fmt.Println("Testing LongestCommonPrefix:")
	for _, test := range tests {
		t.Run(fmt.Sprintf("LongestCommonPrefix(%q)", test.input), func(t *testing.T) {
			result := problem.LongestCommonPrefix(test.input)
			if result != test.expected {
				t.Errorf("expected %q, got %q", test.expected, result)
			}
		})
	}
}
