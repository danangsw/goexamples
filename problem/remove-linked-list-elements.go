package problem

// https://leetcode.com/problems/remove-linked-list-elements/

// Definition for singly-linked list.
/**
type ListNode struct {
	Val int
	Next *ListNode
}
**/

// RemoveElements removes all elements from a linked list that have a specific value.
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}
