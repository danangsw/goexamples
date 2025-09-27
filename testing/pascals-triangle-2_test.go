package testing_test

import (
	"reflect"
	"testing"

	problem "../problem"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{
			name:     "Row 0",
			rowIndex: 0,
			expected: []int{1},
		},
		{
			name:     "Row 1",
			rowIndex: 1,
			expected: []int{1, 1},
		},
		{
			name:     "Row 2",
			rowIndex: 2,
			expected: []int{1, 2, 1},
		},
		{
			name:     "Row 3",
			rowIndex: 3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "Row 4",
			rowIndex: 4,
			expected: []int{1, 4, 6, 4, 1},
		},
		{
			name:     "Row 5",
			rowIndex: 5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
		{
			name:     "Row 10",
			rowIndex: 10,
			expected: []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1},
		},
		{
			name:     "Row 33 (Max Constraint)",
			rowIndex: 33,
			expected: []int{
				1, 33, 528, 5456, 40920, 237336, 1107568, 4272048, 13884156, 38567100, 92561040,
				193536720, 354817320, 573166440, 818809200, 1037158320, 1166803110, 1166803110,
				1037158320, 818809200, 573166440, 354817320, 193536720, 92561040, 38567100,
				13884156, 4272048, 1107568, 237336, 40920, 5456, 528, 33, 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.GetRow(tt.rowIndex)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GetRow(%d) = %v; want %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}
