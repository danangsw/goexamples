package problem

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
func HasPathSum(root *TreeNode, targetSum int) bool {
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
func HasPathSumIterative(root *TreeNode, targetSum int) bool {
	return false
}
