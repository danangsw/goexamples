package testing

import (
	"fmt"
	"testing"

	problem "../problem"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "LeetCode example 1",
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "LeetCode example 2",
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "Three pairs with one single",
			input:    []int{1, 1, 2, 2, 3, 3, 4},
			expected: 4,
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -1, -2},
			expected: -2,
		},
		{
			name:     "Zero as single number",
			input:    []int{0, 1, 1},
			expected: 0,
		},
		{
			name:     "Large numbers",
			input:    []int{100000, 100000, 500000},
			expected: 500000,
		},
		{
			name:     "All pairs except one",
			input:    []int{1, 2, 2, 3, 3, 4, 4, 5, 5},
			expected: 1,
		},
		{
			name:     "Constraint boundary - large array",
			input:    generateLargeArray(),
			expected: 99999,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := problem.SingleNumber(tc.input)
			if result != tc.expected {
				t.Errorf("SingleNumber(%v) = %d; expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

// Helper function for large test case
func generateLargeArray() []int {
	// Create array with pairs 0-99998 and single 99999
	arr := make([]int, 199999) // 99999 pairs + 1 single = 199999 elements
	for i := 0; i < 99999; i++ {
		arr[2*i] = i
		arr[2*i+1] = i
	}
	arr[199998] = 99999 // Single element
	return arr
}

func BenchmarkSingleNumber(b *testing.B) {
	// Benchmark with different array sizes
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			// Create test array with pairs and one single
			arr := make([]int, size*2+1)
			for i := 0; i < size; i++ {
				arr[2*i] = i
				arr[2*i+1] = i
			}
			arr[size*2] = size // Single element

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				problem.SingleNumber(arr)
			}
		})
	}
}
