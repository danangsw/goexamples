package testing_test

import (
	"fmt"
	"reflect"
	"testing"

	"../problem"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Base case n=0",
			n:        0,
			expected: []int{0},
		},
		{
			name:     "Base case n=1",
			n:        1,
			expected: []int{0, 1},
		},
		{
			name:     "Small case n=2",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "Small case n=5",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:     "Medium case n=10",
			n:        10,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
		{
			name:     "Larger case n=15",
			n:        15,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4},
		},
		{
			name:     "Larger case n=20",
			n:        20,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2},
		},
		{
			name:     "Larger case n=25",
			n:        25,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3},
		},
		{
			name:     "Edge Larger case n=105",
			n:        105,
			expected: []int{0,1,1,2,1,2,2,3,1,2,2,3,2,3,3,4,1,2,2,3,2,3,3,4,2,3,3,4,3,4,4,5,1,2,2,3,2,3,3,4,2,3,3,4,3,4,4,5,2,3,3,4,3,4,4,5,3,4,4,5,4,5,5,6,1,2,2,3,2,3,3,4,2,3,3,4,3,4,4,5,2,3,3,4,3,4,4,5,3,4,4,5,4,5,5,6,2,3,3,4,3,4,4,5,3,4},
		},
	}

	// Run tests: Shifted based approach
	fmt.Println("Testing CountBits with Shifted Based Approach:")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.CountBits(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountBits(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}

	// Run tests: Dynamic Programming approach
	fmt.Println("Testing CountBits with Dynamic Programming Approach:")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.CountBitsDP(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountBitsDP(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}

	// Run tests: Modulo-based approach
	fmt.Println("Testing CountBits with Modulo-Based Approach:")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.CountBitsModulo(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountBitsModulo(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}