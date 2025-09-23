package testing_test

import (
	"sort"
	"testing"

	problem "../problem"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, c := range cases {
		got := problem.TwoSum(c.nums, c.target)
		if !equal(got, c.want) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", c.nums, c.target, got, c.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create copies to avoiding mutation
	a1 := make([]int, len(a))
	b1 := make([]int, len(b))

	copy(a1, a)
	copy(b1, b)

	sort.Ints(a1)
	sort.Ints(b1)

	for i := range a1 {
		if a1[i] != b[i] {
			return false
		}
	}

	return true
}
