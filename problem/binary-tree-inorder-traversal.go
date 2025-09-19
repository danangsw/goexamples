// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
package problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InorderIterativeTraversal(root *TreeNode) []int {
	result := []int{}      // Result slice to store inorder traversal
	stack := []*TreeNode{} // Stack to simulate recursion using iteration
	current := root        // Start with the root node

	// Iterate until all nodes are processed
	for current != nil || len(stack) > 0 {
		if current != nil {
			// Go as left as possible
			stack = append(stack, current) // Push current node onto stack
			current = current.Left         // Move to left child
		} else {
			// Backtrack to the last node and process it
			n := len(stack) - 1                  // Index of the last element
			current = stack[n]                   // Pop from stack
			stack = stack[:n]                    // Remove last element from stack
			result = append(result, current.Val) // Add node value to result
			current = current.Right              // Move to right child
		}
	}

	return result
}
