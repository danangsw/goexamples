package testing_test

import (
	"reflect"
	"testing"

	problem "../problem"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		digits []int
		expect []int
	}{
		{
			digits: []int{1, 2, 3},
			expect: []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			expect: []int{4, 3, 2, 2},
		},
		{
			digits: []int{9},
			expect: []int{1, 0},
		},
		{
			digits: []int{9, 9},
			expect: []int{1, 0, 0},
		},
		{
			digits: []int{1, 9},
			expect: []int{2, 0},
		},
		{
			digits: []int{9, 8, 9},
			expect: []int{9, 9, 0},
		},
		{
			digits: []int{9, 9, 9},
			expect: []int{1, 0, 0, 0},
		},
		{
			digits: []int{0},
			expect: []int{1},
		},
		{
			digits: []int{1, 0},
			expect: []int{1, 1},
		},
		{
			digits: []int{1, 0, 0},
			expect: []int{1, 0, 1},
		},
	}

	for _, tc := range testCases {
		result := problem.PlusOne(tc.digits)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("For digits=%v, expected %v, but got %v", tc.digits, tc.expect, result)
		}
	}
}
