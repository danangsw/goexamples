package problem

// https://leetcode.com/problems/balanced-binary-tree/
/**
 * Definition for a binary tree node.
 */
/**
 type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
**/

import "../helper"

func IsBalanced(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func IsBalancedOptimized(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := checkBalance(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := checkBalance(root.Right)
	if rightDepth == -1 {
		return -1
	}
	if helper.Abs(leftDepth-rightDepth) > 1 {
		return -1
	}

	return helper.Max(leftDepth, rightDepth) + 1
}
