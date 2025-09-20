// https://leetcode.com/problems/binary-tree-inorder-traversal/description/
package problem
import (
	helper "../helper"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// InorderIterativeTraversal performs an inorder traversal of a binary tree iteratively (without recursion).
// It returns a slice of integers representing the values of the nodes in inorder sequence.
// Inorder: Left -> Node -> Right (LNR)
// Best For: Copying, Serializing, and Deserializing a tree
// Key Characteristics:Process parent before children
func InorderIterativeTraversal(root *helper.TreeNode) []int {
	result := []int{}      // Result slice to store inorder traversal
	stack := []*helper.TreeNode{} // Stack to simulate recursion using iteration
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

// PreorderIterativeTraversal performs a preorder traversal of a binary tree iteratively (without recursion).
// It returns a slice of integers representing the values of the nodes in preorder sequence.
// Preorder: Node -> Left -> Right (NLR)
// Best For: BST operations like insertion and deletion, sorted output
// Key Charateristics: Gives sorted output for BSTs
func PreorderIterativeTraversal(root *helper.TreeNode) []int {
	result := []int{}      // Result slice to store preorder traversal
	stack := []*helper.TreeNode{} // Stack to simulate recursion using iteration

	if root == nil {
		return result
	}

	stack = append(stack, root) // Start with the root node

	for len(stack) > 0 {
		// Pop from stack
		n := len(stack) - 1                  // Index of the last element
		current := stack[n]                  // Get the last element
		stack = stack[:n]                    // Remove last element from stack
		result = append(result, current.Val) // Add node value to result

		// Push right first
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		// Push left second so that left is processed first
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}
