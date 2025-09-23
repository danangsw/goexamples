package problem

// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path
// from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import helper "../helper"
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MinDepth(root.Left)
	right := MinDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return helper.Min(left, right) + 1
}