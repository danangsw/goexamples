package testing_test

import (
	"testing"
	"../problem"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Example 1: n=19 (happy number)",
			n:        19,
			expected: true,
		},
		{
			name:     "Example 2: n=2 (not a happy number)",
			n:        2,
			expected: false,
		},
		{
			name:     "Single digit happy number: n=7",
			n:        7,
			expected: true,
		},
		{
			name:     "Single digit non-happy number: n=4",
			n:        4,
			expected: false,
		},
		{
			name:     "Large happy number: n=100",
			n:        100,
			expected: true,
		},
		{
			name:     "Large non-happy number: n=123",
			n:        123,
			expected: false,
		},
		{
			name:     "Edge case: n=1 (happy number)",
			n:        1,
			expected: true,
		}, 
		{
			name:     "Edge case: n=0 (not a happy number)",
			n:        0,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.IsHappy(tt.n)
			if result != tt.expected {
				t.Errorf("IsHappy(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}
