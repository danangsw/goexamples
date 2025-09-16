package testing_test

import (
	"testing"
	problem "../problem"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "Both single digit, no carry",
			a:        "0",
			b:        "1",
			expected: "1",
		},
		{
			name:     "Both single digit, with carry",
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			name:     "Different lengths, no carry",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "Different lengths, with carry",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "Empty string a",
			a:        "",
			b:        "101",
			expected: "101",
		},
		{
			name:     "Empty string b",
			a:        "110",
			b:        "",
			expected: "110",
		},
		{
			name:     "Long binary strings",
			a:        "1111",
			b:        "1111",
			expected: "11110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.AddBinary(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("AddBinary(%q, %q) = %q; want %q", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}