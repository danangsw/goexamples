package testing_test

import (
    "reflect"
    "testing"
	problem "../problem"
)

func TestMergeSortedA(t *testing.T) {
    testCases := []struct {
        nums1  []int
        m      int
        nums2  []int
        n      int
        expect []int
    }{
        {
            nums1:  []int{1, 2, 3, 0, 0, 0},
            m:      3,
            nums2:  []int{2, 5, 6},
            n:      3,
            expect: []int{1, 2, 2, 3, 5, 6},
        },
        {
            nums1:  []int{1},
            m:      1,
            nums2:  []int{},
            n:      0,
            expect: []int{1},
        },
        {
            nums1:  []int{0},
            m:      0,
            nums2:  []int{1},
            n:      1,
            expect: []int{1},
        },
        {
            nums1:  []int{4, 5, 6, 0, 0, 0},
            m:      3,
            nums2:  []int{1, 2, 3},
            n:      3,
            expect: []int{1, 2, 3, 4, 5, 6},
        },
        {
            nums1:  []int{1, 2, 3, 0, 0, 0},
            m:      3,
            nums2:  []int{4, 5, 6},
            n:      3,
            expect: []int{1, 2, 3, 4, 5, 6},
        },
        {
            nums1:  []int{0, 0, 0, 0, 0},
            m:      0,
            nums2:  []int{1, 2, 3, 4, 5},
            n:      5,
            expect: []int{1, 2, 3, 4, 5},
        },
        {
            nums1:  []int{1, 2, 3, 4, 5},
            m:      5,
            nums2:  []int{},
            n:      0,
            expect: []int{1, 2, 3, 4, 5},
        },
        {
            nums1:  []int{2, 0},
            m:      1,
            nums2:  []int{1},
            n:      1,
            expect: []int{1, 2},
        },
        {
            nums1:  []int{1, 2, 4, 5, 6, 0},
            m:      5,
            nums2:  []int{3},
            n:      1,
            expect: []int{1, 2, 3, 4, 5, 6},
        },
        {
            nums1:  []int{-4, -2, 0, 0, 0, 0},
            m:      3,
            nums2:  []int{-3, -1, 1},
            n:      3,
            expect: []int{-4, -3, -2, -1, 0, 1},
        },
    }

    for _, tc := range testCases {
        nums1Copy := make([]int, len(tc.nums1))
        copy(nums1Copy, tc.nums1)

        problem.MergeSortedArray(tc.nums1, tc.m, tc.nums2, tc.n)

        if !reflect.DeepEqual(tc.nums1, tc.expect) {
            t.Errorf("For nums1=%v, m=%v, nums2=%v, n=%v, expected %v, but got %v", nums1Copy, tc.m, tc.nums2, tc.n, tc.expect, tc.nums1)
        }
    }
}