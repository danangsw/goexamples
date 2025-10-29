package testing_test

import (
	"testing"

	problem "../problem"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected int
	}{
		{
			name:     "Example 1: 11 (binary 1011)",
			input:    11,
			expected: 3,
		},
		{
			name:     "Example 2: 128 (binary 10000000)",
			input:    128,
			expected: 1,
		},
		{
			name:     "Example 3: 4294967293 (binary 11111111111111111111111111111101)",
			input:    4294967293,
			expected: 31,
		},
		{
			name:     "All bits zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "All bits one",
			input:    4294967295,
			expected: 32,
		},
		{
			name:     "Alternating bits starting with 1",
			input:    2863311530, // binary 10101010101010101010101010101010
			expected: 16,
		},
		{
			name:     "Alternating bits starting with 0",
			input:    1431655765, // binary 01010101010101010101010101010101
			expected: 16,
		},
		{
			name:     "Single bit set at position 0",
			input:    1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.HammingWeight(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
