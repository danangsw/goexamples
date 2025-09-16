package testing

import (
	"fmt"
	"math"
	"testing"

	problem "../problem"
)

func TestSqrtx(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		// Edge cases
		{"Zero", 0, 0},
		{"One", 1, 1},

		// Perfect squares
		{"Perfect square 4", 4, 2},
		{"Perfect square 9", 9, 3},
		{"Perfect square 16", 16, 4},
		{"Perfect square 25", 25, 5},
		{"Perfect square 36", 36, 6},
		{"Perfect square 49", 49, 7},
		{"Perfect square 64", 64, 8},
		{"Perfect square 81", 81, 9},
		{"Perfect square 100", 100, 10},
		{"Perfect square 121", 121, 11},
		{"Perfect square 144", 144, 12},

		// Non-perfect squares (floor of sqrt)
		{"Non-perfect 2", 2, 1},
		{"Non-perfect 3", 3, 1},
		{"Non-perfect 5", 5, 2},
		{"Non-perfect 6", 6, 2},
		{"Non-perfect 7", 7, 2},
		{"Non-perfect 8", 8, 2},
		{"Non-perfect 10", 10, 3},
		{"Non-perfect 15", 15, 3},
		{"Non-perfect 24", 24, 4},
		{"Non-perfect 35", 35, 5},
		{"Non-perfect 48", 48, 6},
		{"Non-perfect 63", 63, 7},
		{"Non-perfect 80", 80, 8},
		{"Non-perfect 99", 99, 9},

		// Larger numbers
		{"Large perfect square 10000", 10000, 100},
		{"Large perfect square 1000000", 1000000, 1000},
		{"Large non-perfect 999999", 999999, 999},
		{"Large non-perfect 123456", 123456, 351},

		// Very large numbers
		{"Very large 2147395600", 2147395600, 46340}, // Near max int32
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.Sqrtx(tt.input)
			if result != tt.expected {
				t.Errorf("Sqrtx(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test against Go's math.Sqrt for verification
func TestSqrtxVsMathSqrt(t *testing.T) {
	testValues := []int{0, 1, 2, 3, 4, 5, 8, 9, 16, 25, 100, 1000, 10000, 123456, 999999}

	for _, x := range testValues {
		result := problem.Sqrtx(x)
		expected := int(math.Sqrt(float64(x)))

		if result != expected {
			t.Errorf("Sqrtx(%d) = %d; math.Sqrt floor = %d", x, result, expected)
		}
	}
}

// Benchmark the function
func BenchmarkSqrtx(b *testing.B) {
	testCases := []int{1, 4, 9, 16, 25, 100, 1000, 10000, 123456, 999999}

	for _, x := range testCases {
		b.Run(fmt.Sprintf("x=%d", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				problem.Sqrtx(x)
			}
		})
	}
}

// Test precision and accuracy
func TestSqrtxAccuracy(t *testing.T) {
	// Test that our result squared is close to the input
	testValues := []int{2, 3, 5, 6, 7, 8, 10, 15, 24, 35, 48, 63, 80, 99}

	for _, x := range testValues {
		result := problem.Sqrtx(x)

		// Check that result^2 <= x < (result+1)^2
		if result*result > x {
			t.Errorf("Sqrtx(%d) = %d; %d^2 = %d > %d", x, result, result, result*result, x)
		}

		nextSquare := (result + 1) * (result + 1)
		if x >= nextSquare {
			t.Errorf("Sqrtx(%d) = %d; %d >= (%d+1)^2 = %d", x, result, x, result, nextSquare)
		}
	}
}

// Test edge cases and error conditions
func TestSqrtxEdgeCases(t *testing.T) {
	// Test maximum safe integer for this implementation
	maxSafeInt := 2147395600 // sqrt ≈ 46340
	result := problem.Sqrtx(maxSafeInt)
	expected := 46340

	if result != expected {
		t.Errorf("Sqrtx(%d) = %d; expected %d", maxSafeInt, result, expected)
	}
}
