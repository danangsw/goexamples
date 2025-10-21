package testing_test

import (
	"testing"

	"../problem"
)

func TestTitleToNumber(t *testing.T) {
	tests := testCaseGenerator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.TitleToNumber(tt.input)
			if result != tt.expected {
				t.Errorf("TitleToNumber(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTitleToNumberRecursive(t *testing.T) {
	tests := testCaseGenerator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.TitleToNumberRecursive(tt.input)
			if result != tt.expected {
				t.Errorf("TitleToNumber(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func testCaseGenerator() []struct {
	name     string
	input    string
	expected int
}{
	return []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single letter A",
			input:    "A",
			expected: 1,
		},
		{
			name:     "Single letter Z",
			input:    "Z",
			expected: 26,
		},
		{
			name:     "Two letters AA",
			input:    "AA",
			expected: 27,
		},
		{
			name:     "Two letters AB",
			input:    "AB",
			expected: 28,
		},
		{
			name:     "Three letters AAA",
			input:    "AAA",
			expected: 703,
		},
		{
			name:     "Three letters ZZZ",
			input:    "ZZZ",
			expected: 18278,
		},
		{
			name:     "Mixed letters AZ",
			input:    "AZ",
			expected: 52,
		},
		{
			name:     "Edge case latest letters FXSHRXW",
			input:    "FXSHRXW",
			expected: 2147483647,
		},
		{
			name:     "Edge case empty string",
			input:    "",
			expected: 0,
		},
	}
}
