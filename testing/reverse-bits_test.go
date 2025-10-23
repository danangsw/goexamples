package testing_test

import (
	"fmt"
	"testing"

	"../problem"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    43261596,
			expected: 964176192,
		},
		{
			name:     "Example 2",
			input:    4294967293,
			expected: 3221225471,
		},
		{
			name:     "All bits zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "All bits one",
			input:    4294967295,
			expected: 4294967295,
		},
		{
			name:     "Alternating bits starting with 1",
			input:    2863311530,
			expected: 1431655765,
		},
		{
			name:     "Alternating bits starting with 0",
			input:    1431655765,
			expected: 2863311530,
		},
		{
			name:     "Single bit set at LSB",
			input:    1,
			expected: 2147483648,
		},
		{
			name:     "Single bit set at MSB",
			input:    2147483648,
			expected: 1,
		},
	}

	fmt.Println("Running tests for ReverseBits: Table Lookup...")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.ReverseBits(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}

	fmt.Println("Running tests for ReverseBits: iterative bit-reverse (32 shifts)...")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.ReverseBitsIterative(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
