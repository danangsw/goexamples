package problem

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	start := 0
	end := len(nums) - 1
	mid := start + (end - start) / 2

	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[start:mid])
	root.Right = SortedArrayToBST(nums[mid+1 : end+1])
	return root
}
