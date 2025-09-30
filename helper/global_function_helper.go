package helper

// Helper function to create a binary tree from array representation
// nil values in the array represent null nodes
func CreateTreeFromArray(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// Helper function to create int pointer
func IntPtr(val int) *int {
	return &val
}

// Helper function to compare slices
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of an integer
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper function to calculate maximum depth for property testing
func GetMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := GetMaxDepth(root.Left)
	right := GetMaxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// RECOMMENDED (DOCUMENTED):
// BuildLinkedList creates a linked list from a slice of integers.
// The list is constructed in the same order as the input slice.
// Returns nil if the input slice is empty.
// Time Complexity: O(n), Space Complexity: O(n)
func BuildLinkedList(values []int) *ListNode {
	// CURRENT: Only handles empty slice
	if len(values) == 0 {
		return nil
	}

	// COULD ADD: Handle nil slice (though unlikely in practice)
	if values == nil {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

func ListToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
