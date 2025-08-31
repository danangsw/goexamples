package testing

import (
	"fmt"
	"testing"
	problem "../problem"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("RomanToInt(%q)", test.input), func(t *testing.T) {
			result := problem.RomanToInt(test.input)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
