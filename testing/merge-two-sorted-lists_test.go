package testing_test

import (
	"reflect"
	"testing"

	problem "../problem"
)

type ListNode = problem.ListNode

func TestMergeTwoList(t *testing.T) {
	// Helper function to create a linked list from a slice of integers
	createList := func(nums []int) *ListNode {
		if len(nums) == 0 {
			return nil
		}
		head := &ListNode{Val: nums[0]}
		curr := head
		for i := 1; i < len(nums); i++ {
			curr.Next = &ListNode{Val: nums[i]}
			curr = curr.Next
		}
		return head
	}

	// Helper function to convert a linked list to a slice of integers for comparison
	listToSlice := func(head *ListNode) []int {
		result := []int{}
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}
		return result
	}

	testCases := []struct {
		l1     []int
		l2     []int
		expect []int
	}{
		{
			l1:     []int{1, 2, 4},
			l2:     []int{1, 3, 4},
			expect: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:     []int{},
			l2:     []int{},
			expect: []int{},
		},
		{
			l1:     []int{0},
			l2:     []int{0},
			expect: []int{0, 0},
		},
		{
			l1:     []int{},
			l2:     []int{0},
			expect: []int{0},
		},
		{
			l1:     []int{0},
			l2:     []int{},
			expect: []int{0},
		},
		{
			l1:     []int{-2, -1, 0, 1, 2},
			l2:     []int{-3, -2, 2, 3, 4},
			expect: []int{-3, -2, -2, -1, 0, 1, 2, 2, 3, 4},
		},
		{
			l1:     []int{5},
			l2:     []int{1, 2, 3, 4},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			l1:     []int{1, 2, 3, 4},
			l2:     []int{5},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			l1:     []int{1, 1, 1, 1},
			l2:     []int{1, 1, 1, 1},
			expect: []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			l1:     []int{-100, 0, 100},
			l2:     []int{-100, 0, 100},
			expect: []int{-100, -100, 0, 0, 100, 100},
		},
		{
			l1:     []int{-5, -3, -1, 1, 3, 5},
			l2:     []int{-4, -2, 0, 2, 4, 6},
			expect: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6},
		},
		{
			l1:     []int{},
			l2:     []int{-4, -2, 0, 2, 4, 6},
			expect: []int{-4, -2, 0, 2, 4, 6},
		},
		{
			l1:     []int{-5, -3, -1, 1, 3, 5},
			l2:     []int{},
			expect: []int{-5, -3, -1, 1, 3, 5},
		},
	}

	for _, tc := range testCases {
		l1 := createList(tc.l1)
		l2 := createList(tc.l2)
		mergedList := problem.MergeTwoList(l1, l2)
		result := listToSlice(mergedList)

		if !reflect.DeepEqual(result, tc.expect) {
			t.Errorf("For l1=%v and l2=%v, expected %v, but got %v", tc.l1, tc.l2, tc.expect, result)
		}
	}
}
