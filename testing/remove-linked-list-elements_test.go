package testing_test

import (
	"reflect"
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		val      int
		expected []int
	}{
		{
			name:     "Empty list",
			input:    []int{},
			val:      1,
			expected: []int{},
		},
		{
			name:     "Single node to remove",
			input:    []int{1},
			val:      1,
			expected: []int{},
		},
		{
			name:     "Single node to keep",
			input:    []int{2},
			val:      1,
			expected: []int{2},
		},
		{
			name:     "Remove all occurrences",	
			input:    []int{1, 1, 1, 1},
			val:      1,
			expected: []int{},
		},
		{
			name:     "No elements to remove",
			input:    []int{1, 2, 3, 4, 5},
			val:      6,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Remove some elements",
			input:    []int{1, 2, 6, 3, 4, 5, 6},
			val:      6,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Remove head elements",
			input:    []int{7, 7, 7, 8, 9},
			val:      7,
			expected: []int{8, 9},
		},
		{
			name:     "Remove tail elements",
			input:    []int{1, 2, 3, 4, 5, 5, 5},
			val:      5,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Alternate elements to remove",	
			input:    []int{1, 2, 1, 2, 1, 2},
			val:      1,
			expected: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := helper.CreateList(tt.input)
			resultHead := problem.RemoveElements(head, tt.val)
			resultSlice := helper.ListToSlice(resultHead)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("RemoveElements(%v, %d) = %v, expected %v", tt.input, tt.val, resultSlice, tt.expected)
			}
		})
	}
}
