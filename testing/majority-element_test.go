package testing_test

import (
	"testing"

	problem "../problem"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 2, 2}, 1},
	}

	for _, test := range tests {
		result := problem.MajorityElement(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
