package testing

import (
	"testing"

	"../problem"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{nums: []int{1, 3, 5, 6}, target: 0, expect: 0},
		{nums: []int{1, 3, 5, 6}, target: 7, expect: 4},
		{nums: []int{}, target: 5, expect: 0},
		{nums: []int{5}, target: 2, expect: 0},
		{nums: []int{5}, target: 7, expect: 1},
		{nums: []int{5}, target: 5, expect: 0},
		{nums: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{nums: []int{1, 2, 5, 6}, target: 3, expect: 2},
		{nums: []int{1, 3, 4, 5, 6}, target: 2, expect: 1},
		{nums: []int{1, 3, 4, 5, 6}, target: 3, expect: 1},
		{nums: []int{1, 3, 4, 5, 7}, target: 6, expect: 4},
		{nums: []int{-3, -1, 1, 5}, target: 0, expect: 2},
		{nums: []int{1000, 2000, 3000}, target: 2500, expect: 2},
		{nums: []int{-5, -3, 0, 2, 4}, target: -4, expect: 1},
	}

	for _, tc := range testCases {
		result := problem.SearchInsert(tc.nums, tc.target)
		if result != tc.expect {
			t.Errorf("For nums=%v, target=%v, expected %v, but got %v", tc.nums, tc.target, tc.expect, result)
		}
	}
}
