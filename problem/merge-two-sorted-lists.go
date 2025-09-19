// https://leetcode.com/problems/merge-two-sorted-lists

package problem

func MergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	node := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next, l1 = l1, l1.Next
		} else {
			node.Next, l2 = l2, l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return head.Next
}
