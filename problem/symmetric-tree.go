// https://leetcode.com/problems/symmetric-tree
package problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// IsSymmetric determines if a binary tree is symmetric around its center.
// A tree is symmetric if the left subtree is a mirror reflection of the right subtree.
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (for the stack)
//
// Algorithm:
// 1. Use a stack to store pairs of nodes that should be mirror images
// 2. For symmetry, compare left.left with right.right and left.right with right.left
// 3. Values must match and structure must be identical but mirrored
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// Stack to hold pairs of nodes to compare
	stack := [][2]*TreeNode{}
	stack = append(stack, [2]*TreeNode{root.Left, root.Right})

	for len(stack) > 0 {
		// Pop pair from stack
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left, right := pair[0], pair[1]

		// Both nil - symmetric so far, continue
		if left == nil && right == nil {
			continue
		}

		// One nil, one not - not symmetric
		if left == nil || right == nil {
			return false
		}

		// Different values - not symmetric
		if left.Val != right.Val {
			return false
		}

		// For symmetry: left.outer should match right.outer, left.inner should match right.inner
		stack = append(stack, [2]*TreeNode{left.Right, right.Left}) // Inner children
		stack = append(stack, [2]*TreeNode{left.Left, right.Right}) // Outer children
	}

	return true
}
