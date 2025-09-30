package problem

// https://leetcode.com/problems/linked-list-cycle/

import (
	"../helper"
)

// Definition for singly-linked list.
type ListNode = helper.ListNode

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
