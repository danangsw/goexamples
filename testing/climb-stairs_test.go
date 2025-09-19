package testing_test

import (
	"testing"

	"../problem"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Base case n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "Base case n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "Small case n=3",
			n:        3,
			expected: 3,
		},
		{
			name:     "Medium case n=4",
			n:        4,
			expected: 5,
		},
		{
			name:     "Medium case n=5",
			n:        5,
			expected: 8,
		},
		{
			name:     "Larger case n=10",
			n:        10,
			expected: 89,
		},
		{
			name:     "Larger case n=15",
			n:        15,
			expected: 987,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.ClimbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("ClimbStairs(%d) = %d, expected %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestClimbStairsEdgeCases(t *testing.T) {
	// Test edge cases separately for clarity
	edgeCases := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Minimum valid input n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "Second minimum n=2",
			n:        2,
			expected: 2,
		},
	}

	for _, tt := range edgeCases {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.ClimbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("ClimbStairs(%d) = %d, expected %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestClimbStairsFibonacciSequence(t *testing.T) {
	// Verify that the function follows Fibonacci sequence
	// F(1)=1, F(2)=2, F(n)=F(n-1)+F(n-2)
	fibonacciExpected := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	for i, expected := range fibonacciExpected {
		n := i + 1
		result := problem.ClimbStairs(n)
		if result != expected {
			t.Errorf("ClimbStairs(%d) = %d, expected %d (Fibonacci sequence verification)", n, result, expected)
		}
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"Small_n5", 5},
		{"Medium_n20", 20},
		{"Large_n40", 40},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				problem.ClimbStairs(bm.n)
			}
		})
	}
}
