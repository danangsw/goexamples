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

func MinDepthIterativeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		level := len(queue)

		// Process all nodes at the current level
		for i := 0; i < level; i++ {
			node := queue[0]
			queue = queue[1:]

			// Check if it's a leaf node - the minimum depth is found
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// add children to the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}	
		}
	}

	return depth
}
