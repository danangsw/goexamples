package testing_test

import (
	"reflect"
	"testing"

	problem "../problem"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
        nums   []int
        expect int
        result []int
    }{
        {
            nums:   []int{1, 1, 2},
            expect: 2,
            result: []int{1, 2},
        },
        {
            nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
            expect: 5,
            result: []int{0, 1, 2, 3, 4},
        },
        {
            nums:   []int{},
            expect: 0,
            result: []int{},
        },
        {
            nums:   []int{1},
            expect: 1,
            result: []int{1},
        },
        {
            nums:   []int{-100, -100, -99, -99, -98},
            expect: 3,
            result: []int{-100, -99, -98},
        },
        {
            nums:   []int{-1, 0, 0, 0, 3, 3},
            expect: 3,
            result: []int{-1, 0, 3},
        },
        {
            nums:   []int{1, 1, 1, 1, 1},
            expect: 1,
            result: []int{1},
        },
        {
            nums:   []int{1, 2, 3, 4, 5},
            expect: 5,
            result: []int{1, 2, 3, 4, 5},
        },
        {
            nums:   []int{-100, 100},
            expect: 2,
            result: []int{-100, 100},
        },
        {
            nums:   []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
            expect: 11,
            result: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
        },
    }

	for _, tc := range testCases {
		numsCopy := make([]int, len(tc.nums))
		copy(numsCopy, tc.nums)

		k := problem.RemoveDuplicates(tc.nums)

		if k != tc.expect {
			t.Errorf("For nums=%v, expected length %v, but got %v", numsCopy, tc.expect, k)
		}

		if !reflect.DeepEqual(tc.nums[:k], tc.result) && len(tc.result) > 0 {
			t.Errorf("For nums=%v, expected result %v, but got %v", numsCopy, tc.result, tc.nums[:k])
		}
	}
}
