package problem

import helper "../helper"

// https://leetcode.com/problems/path-sum/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// HasPathSum checks if the tree has a root-to-leaf path such that adding up all
// the values along the path equals the given sum using a recursive approach.
func HasPathSum(root *helper.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	targetSum -= root.Val
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}

// HasPathSumIterative checks if the tree has a root-to-leaf path such that adding up
// all the values along the path equals the given sum using an iterative approach.
func HasPathSumIterative(root *helper.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// stack simulates the call stack for DFS
	stack := []struct {
		node *helper.TreeNode
		sum  int
	}{{root, targetSum - root.Val}}

	for len(stack) > 0 {
		// Pop the top element
		n := len(stack) - 1
		current := stack[n]
		stack = stack[:n]

		// If it's a leaf node and the remaining sum is zero, we found a path
		if current.node.Left == nil && current.node.Right == nil && current.sum == 0 {
			return true
		}

		// if it has a right chilld, push right child and updated sum onto stack
		if current.node.Right != nil {
			stack = append(stack, struct {
				node *helper.TreeNode
				sum  int
			}{current.node.Right, current.sum - current.node.Right.Val})
		}

		// if it has a left chilld, push left child and updated sum onto stack
		if current.node.Left != nil {
			stack = append(stack, struct {
				node *helper.TreeNode
				sum  int
			}{current.node.Left, current.sum - current.node.Left.Val})
		}
	}

	return false
}
