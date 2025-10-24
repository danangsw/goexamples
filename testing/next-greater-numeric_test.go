package testing_test

import (
	"fmt"
	"testing"

	"../problem"
)

func TestNextGreaterNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    12,
			expected: 22,
		},
		{
			name:     "Example 2",
			input:    21,
			expected: 22,
		},
		{
			name:     "All digits same",
			input:    111,
			expected: 122,
		},
		{
			name:     "Descending order digits",
			input:    321,
			expected: 333,
		},
		{
			name:     "Ascending order digits",
			input:    1234,
			expected: 1333,
		},
		{
			name:     "Large number with mixed digits",
			input:    534976,
			expected: 551555,
		},
		{
			name:     "Number with trailing zeros",
			input:    1200,
			expected: 1333,
		},
		{
			name:     "Single digit number",
			input:    7,
			expected: 22,
		},
		{
			name:     "Maximum 32-bit integer case",
			input:    2147483647,
			expected: 2172777777,
		},
		{
			name:     "Number resulting in overflow",
			input:    1999999999,
			expected: 2123334444,
		},
		{
			name:     "Number with zeros in between",
			input:    10203,
			expected: 14444,
		},
		{
			name:     "Number 1",
			input:    1,
			expected: 22,
		},
		{
			name:     "Number 1000",
			input:    1000,
			expected: 1333,
		},
		{
			name:     "Number 3000",
			input:    3000,
			expected: 3133,
		},
		{
			name:     "Number 1551",
			input:    1551,
			expected: 3133,
		},
	}
	fmt.Println("Running TestNextGreaterNumeric: Brute force approach")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.NextGreaterNumber(tt.input)
			if result != tt.expected {
				t.Errorf("Brute force: NextGreaterNumber(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}

	fmt.Println("Running TestNextGreaterNumeric: Precomputed approach")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.NextGreaterNumberPrecomputed(tt.input)
			if result != tt.expected {
				t.Errorf("Precomputed: NextGreaterNumber(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
