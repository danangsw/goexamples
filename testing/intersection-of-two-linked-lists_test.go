package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestGetIntersectionNode(t *testing.T) {
	cases := []struct {
		name        string
		listA       []int
		listB       []int
		intersectAt int // -1 means no intersection
		expected    int // -1 means nil
	}{
		{
			name:        "No intersection",
			listA:       []int{1, 2, 3},
			listB:       []int{4, 5, 6},
			intersectAt: -1,
			expected:    -1,
		},
		{
			name:        "Intersection at middle",
			listA:       []int{1, 2, 3, 4, 5},
			listB:       []int{9, 8, 3, 4, 5}, // Intersects at 3
			intersectAt: 2,
			expected:    3,
		},
		{
			name:        "Intersection at head",
			listA:       []int{1, 2, 3},
			listB:       []int{1, 2, 3}, // Same list
			intersectAt: 0,
			expected:    1,
		},
		// Add more comprehensive cases...
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create lists and set up intersection
			headA := helper.BuildLinkedList(tc.listA)
			headB := helper.BuildLinkedList(tc.listB)

			if tc.intersectAt != -1 {
				// Find intersection point and link lists
				intersectNode := findNodeAtIndex(headA, tc.intersectAt)
				findLastNode(headB).Next = intersectNode
			}

			result := problem.GetIntersectionNode(headA, headB)

			if tc.expected == -1 {
				if result != nil {
					t.Errorf("Expected nil, got %v", result.Val)
				}
			} else {
				if result == nil || result.Val != tc.expected {
					t.Errorf("Expected %d, got %v", tc.expected, result.Val)
				}
			}
		})
	}
}

// Helper functions for test setup
func findNodeAtIndex(head *helper.ListNode, index int) *helper.ListNode {
	current := head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	return current
}

func findLastNode(head *helper.ListNode) *helper.ListNode {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	return current
}
