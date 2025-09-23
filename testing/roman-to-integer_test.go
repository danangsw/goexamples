package testing_test

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

	fmt.Println("Testing RomanToIntRL:")
	for _, test := range tests {
		t.Run(fmt.Sprintf("RomanToIntRL(%q)", test.input), func(t *testing.T) {
			result := problem.RomanToIntRL(test.input)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}

	fmt.Println("Testing RomanToIntLR:")
	for _, test := range tests {
		t.Run(fmt.Sprintf("RomanToIntLR(%q)", test.input), func(t *testing.T) {
			result := problem.RomanToIntLR(test.input)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
