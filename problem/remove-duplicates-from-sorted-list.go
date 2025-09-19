// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
// Given the head of a sorted linked list,
// delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
package problem

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current != nil {
		// If the next node is a duplicate, skip it
		if current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next // Skip the duplicate
		} else { // Move to the next node
			current = current.Next
		}
	}
	return head
}
