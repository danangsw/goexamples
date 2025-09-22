package testing_test

import (
	"math"
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool // We'll test if the result is a valid height-balanced BST
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Single element",
			input:    []int{0},
			expected: true,
		},
		{
			name:     "Two elements",
			input:    []int{1, 3},
			expected: true,
		},
		{
			name:     "Three elements",
			input:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "LeetCode Example 1: [-10,-3,0,5,9]",
			input:    []int{-10, -3, 0, 5, 9},
			expected: true,
		},
		{
			name:     "LeetCode Example 2: [1,3]",
			input:    []int{1, 3},
			expected: true,
		},
		{
			name:     "Four elements",
			input:    []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "Five elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Six elements",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: true,
		},
		{
			name:     "Seven elements (complete binary tree)",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "Negative numbers",
			input:    []int{-5, -3, -1, 0, 2, 4, 6},
			expected: true,
		},
		{
			name:     "All negative",
			input:    []int{-10, -8, -6, -4, -2},
			expected: true,
		},
		{
			name:     "All positive",
			input:    []int{1, 3, 5, 7, 9, 11, 13},
			expected: true,
		},
		{
			name:     "Large array (15 elements)",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: true,
		},
		{
			name:     "Consecutive numbers",
			input:    []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: true,
		},
		{
			name:     "Large gaps",
			input:    []int{1, 100, 200, 300, 400},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.SortedArrayToBST(tt.input)

			// Check if result is nil for empty input
			if len(tt.input) == 0 {
				if result != nil {
					t.Errorf("Expected nil for empty input, got non-nil")
				}
				return
			}

			// Verify the tree is a valid BST
			if !isValidBST(result) {
				t.Errorf("Result is not a valid BST")
			}

			// Verify the tree is height-balanced
			if !isHeightBalanced(result) {
				t.Errorf("Result is not height-balanced")
			}

			// Verify all elements from input are present in the tree
			treeElements := getInorderTraversal(result)
			if !helper.SlicesEqual(treeElements, tt.input) {
				t.Errorf("Tree elements %v don't match input %v", treeElements, tt.input)
			}

			// Verify tree size matches input size
			treeSize := getTreeSize(result)
			if treeSize != len(tt.input) {
				t.Errorf("Tree size %d doesn't match input size %d", treeSize, len(tt.input))
			}
		})
	}
}

func TestSortedArrayToBSTEdgeCases(t *testing.T) {
	t.Run("Nil input should be handled gracefully", func(t *testing.T) {
		// Note: Go slices are different from nil, but empty slice should work
		result := problem.SortedArrayToBST([]int{})
		if result != nil {
			t.Errorf("Expected nil for empty slice, got %v", result)
		}
	})

	t.Run("Single element creates single node", func(t *testing.T) {
		result := problem.SortedArrayToBST([]int{42})

		if result == nil {
			t.Fatal("Expected non-nil result for single element")
		}

		if result.Val != 42 {
			t.Errorf("Expected root value 42, got %d", result.Val)
		}

		if result.Left != nil || result.Right != nil {
			t.Errorf("Expected leaf node, but has children")
		}
	})

	t.Run("Two elements create proper parent-child relationship", func(t *testing.T) {
		result := problem.SortedArrayToBST([]int{1, 2})

		if result == nil {
			t.Fatal("Expected non-nil result")
		}

		// Either 1 or 2 can be root (both valid)
		if result.Val == 1 {
			// Root is 1, right child should be 2
			if result.Left != nil {
				t.Errorf("Expected no left child when root is 1")
			}
			if result.Right == nil || result.Right.Val != 2 {
				t.Errorf("Expected right child to be 2")
			}
		} else if result.Val == 2 {
			// Root is 2, left child should be 1
			if result.Right != nil {
				t.Errorf("Expected no right child when root is 2")
			}
			if result.Left == nil || result.Left.Val != 1 {
				t.Errorf("Expected left child to be 1")
			}
		} else {
			t.Errorf("Root should be either 1 or 2, got %d", result.Val)
		}
	})

	t.Run("Large input creates balanced tree", func(t *testing.T) {
		// Create array of 31 elements (perfect binary tree)
		input := make([]int, 31)
		for i := 0; i < 31; i++ {
			input[i] = i + 1
		}

		result := problem.SortedArrayToBST(input)

		if result == nil {
			t.Fatal("Expected non-nil result for large input")
		}

		height := getTreeHeight(result)
		expectedHeight := 5 // log2(31) + 1 ≈ 5

		if height != expectedHeight {
			t.Errorf("Expected height %d for 31 elements, got %d", expectedHeight, height)
		}

		if !isHeightBalanced(result) {
			t.Errorf("Large tree should be height-balanced")
		}
	})

	t.Run("Negative and positive mixed", func(t *testing.T) {
		input := []int{-100, -50, 0, 25, 75, 100}
		result := problem.SortedArrayToBST(input)

		if !isValidBST(result) {
			t.Errorf("Mixed negative/positive tree should be valid BST")
		}

		if !isHeightBalanced(result) {
			t.Errorf("Mixed negative/positive tree should be height-balanced")
		}
	})
}

func TestSortedArrayToBSTProperties(t *testing.T) {
	t.Run("BST property maintained", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 4, 5},
			{-5, -2, 0, 3, 7, 10},
			{10, 20, 30, 40, 50, 60, 70},
		}

		for _, input := range testCases {
			result := problem.SortedArrayToBST(input)

			if !isValidBST(result) {
				t.Errorf("BST property violated for input %v", input)
			}
		}
	})

	t.Run("Height balance property maintained", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3},
			{1, 2, 3, 4},
			{1, 2, 3, 4, 5, 6, 7},
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		}

		for _, input := range testCases {
			result := problem.SortedArrayToBST(input)

			if !isHeightBalanced(result) {
				t.Errorf("Height balance property violated for input %v", input)
				t.Logf("Tree height: %d", getTreeHeight(result))
			}
		}
	})

	t.Run("All elements preserved", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 9, 11, 13, 15}
		result := problem.SortedArrayToBST(input)

		treeElements := getInorderTraversal(result)

		if !helper.SlicesEqual(treeElements, input) {
			t.Errorf("Elements not preserved. Expected %v, got %v", input, treeElements)
		}
	})

	t.Run("Optimal tree height", func(t *testing.T) {
		// Test that height is approximately log2(n)
		sizes := []int{7, 15, 31, 63}
		expectedHeights := []int{3, 4, 5, 6}

		for i, size := range sizes {
			input := make([]int, size)
			for j := 0; j < size; j++ {
				input[j] = j + 1
			}

			result := problem.SortedArrayToBST(input)
			height := getTreeHeight(result)

			if height != expectedHeights[i] {
				t.Errorf("For size %d, expected height %d, got %d",
					size, expectedHeights[i], height)
			}
		}
	})
}

func BenchmarkSortedArrayToBST(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"SortedArrayToBST_Small_7", 7},
		{"SortedArrayToBST_Medium_31", 31},
		{"SortedArrayToBST_Large_127", 127},
		{"SortedArrayToBST_XLarge_511", 511},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Create sorted array
			input := make([]int, bm.size)
			for i := 0; i < bm.size; i++ {
				input[i] = i + 1
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				problem.SortedArrayToBST(input)
			}
		})
	}
}

/***
func ExampleSortedArrayToBST() {
	// Create BST from sorted array
	nums := []int{-10, -3, 0, 5, 9}
	root := problem.SortedArrayToBST(nums)

	// Verify it's a valid BST by doing inorder traversal
	elements := getInorderTraversal(root)

	// Print the elements (should be same as input)
	for i, val := range elements {
		if i > 0 {
			print(", ")
		}
		print(val)
	}
	println()
	// Output: -10, -3, 0, 5, 9
}
***/

// Helper function to check if tree is a valid BST
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}

	return isValidBSTHelper(node.Left, min, int64(node.Val)) &&
		isValidBSTHelper(node.Right, int64(node.Val), max)
}

// Helper function to check if tree is height-balanced
func isHeightBalanced(root *TreeNode) bool {
	_, balanced := checkHeight(root)
	return balanced
}

func checkHeight(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkHeight(node.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := checkHeight(node.Right)
	if !rightBalanced {
		return 0, false
	}

	heightDiff := leftHeight - rightHeight
	if heightDiff < 0 {
		heightDiff = -heightDiff
	}

	balanced := heightDiff <= 1
	height := max(leftHeight, rightHeight) + 1

	return height, balanced
}

// Helper function to get inorder traversal
func getInorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, getInorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, getInorderTraversal(root.Right)...)

	return result
}

// Helper function to get tree size
func getTreeSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getTreeSize(root.Left) + getTreeSize(root.Right)
}

// Helper function to get tree height
func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getTreeHeight(root.Left)
	rightHeight := getTreeHeight(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

// Helper function for max (for Go versions < 1.21)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
