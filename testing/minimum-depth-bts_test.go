package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected int
	}{
		{
			name:     "Empty tree",
			input:    []*int{},
			expected: 0,
		},
		{
			name:     "Single node",
			input:    []*int{helper.IntPtr(1)},
			expected: 1,
		},
		{
			name:     "Two nodes - left child only",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2)},
			expected: 2,
		},
		{
			name:     "Two nodes - right child only",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2)},
			expected: 2,
		},
		{
			name:     "LeetCode Example 1: [3,9,20,null,null,15,7]",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(9), helper.IntPtr(20), nil, nil, helper.IntPtr(15), helper.IntPtr(7)},
			expected: 2, // Path: 3 → 9 (leaf)
		},
		{
			name:     "LeetCode Example 2: [2,null,3,null,4,null,5,null,6]",
			input:    []*int{helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4), nil, helper.IntPtr(5), nil, helper.IntPtr(6)},
			expected: 5, // Linear right-skewed tree
		},
		{
			name:     "Perfect binary tree [1,2,3,4,5,6,7]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), helper.IntPtr(6), helper.IntPtr(7)},
			expected: 3, // All leaves at level 3
		},
		{
			name:     "Left-skewed tree [1,2,null,3,null,4]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4)},
			expected: 4, // Path: 1 → 2 → 3 → 4
		},
		{
			name:     "Right-skewed tree [1,null,2,null,3,null,4]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4)},
			expected: 4, // Path: 1 → 2 → 3 → 4
		},
		{
			name:     "Unbalanced tree with early leaf [1,2,3,4,null,null,5]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), nil, nil, helper.IntPtr(5)},
			expected: 3, // Path: 1 → 3 → 5
		},
		{
			name:     "Tree with nulls [1,2,3,null,4,null,5]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), nil, helper.IntPtr(4), nil, helper.IntPtr(5)},
			expected: 3, // Paths: 1 → 2 → 4 or 1 → 3 → 5
		},
		{
			name:     "Deep left, shallow right [1,2,null,3,null,4,null,5]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4), nil, helper.IntPtr(5)},
			expected: 5, // Only one path: 1 → 2 → 3 → 4 → 5
		},
		{
			name:     "Negative values [-1,-2,-3]",
			input:    []*int{helper.IntPtr(-1), helper.IntPtr(-2), helper.IntPtr(-3)},
			expected: 2, // Both -2 and -3 are leaves at depth 2
		},
		{
			name:     "Mixed positive and negative [0,-1,1,-2,null,null,2]",
			input:    []*int{helper.IntPtr(0), helper.IntPtr(-1), helper.IntPtr(1), helper.IntPtr(-2), nil, nil, helper.IntPtr(2)},
			expected: 3, // Paths: 0 → -1 → -2 or 0 → 1 → 2
		},
		{
			name:     "Large values [100,50,150,25,75,125,175]",
			input:    []*int{helper.IntPtr(100), helper.IntPtr(50), helper.IntPtr(150), helper.IntPtr(25), helper.IntPtr(75), helper.IntPtr(125), helper.IntPtr(175)},
			expected: 3, // Perfect binary tree, all leaves at level 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.MinDepth(root)

			if result != tt.expected {
				t.Errorf("MinDepth() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinDepthEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := problem.MinDepth(nil)
		expected := 0

		if result != expected {
			t.Errorf("MinDepth(nil) = %d, expected %d", result, expected)
		}
	})

	t.Run("Single node with zero value", func(t *testing.T) {
		root := &TreeNode{Val: 0}
		result := problem.MinDepth(root)
		expected := 1

		if result != expected {
			t.Errorf("MinDepth with zero value = %d, expected %d", result, expected)
		}
	})

	t.Run("Deep left-only chain", func(t *testing.T) {
		// Create: 1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Left = &TreeNode{Val: i}
			current = current.Left
		}

		result := problem.MinDepth(root)
		expected := 10

		if result != expected {
			t.Errorf("Deep left chain = %d, expected %d", result, expected)
		}
	})

	t.Run("Deep right-only chain", func(t *testing.T) {
		// Create: 1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
		}

		result := problem.MinDepth(root)
		expected := 10

		if result != expected {
			t.Errorf("Deep right chain = %d, expected %d", result, expected)
		}
	})

	t.Run("Zigzag pattern", func(t *testing.T) {
		// Create a zigzag pattern:
		//     1
		//    /
		//   2
		//    \
		//     3
		//    /
		//   4
		//    \
		//     5 (leaf)
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 3}
		root.Left.Right.Left = &TreeNode{Val: 4}
		root.Left.Right.Left.Right = &TreeNode{Val: 5}

		result := problem.MinDepth(root)
		expected := 5

		if result != expected {
			t.Errorf("Zigzag pattern = %d, expected %d", result, expected)
		}
	})

	t.Run("Multiple paths with different depths", func(t *testing.T) {
		// Create tree with multiple leaf depths:
		//       1
		//      / \
		//     2   3  ← leaf at depth 2
		//    /
		//   4    ← leaf at depth 3
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3} // This is a leaf at depth 2
		root.Left.Left = &TreeNode{Val: 4} // This is a leaf at depth 3

		result := problem.MinDepth(root)
		expected := 2 // Should find the shallower leaf (node 3)

		if result != expected {
			t.Errorf("Multiple paths = %d, expected %d", result, expected)
		}
	})

	t.Run("Complex unbalanced tree", func(t *testing.T) {
		// Create a complex tree:
		//         1
		//       /   \
		//      2     3
		//     / \     \
		//    4   5     6  ← leaf at depth 3
		//   /   /
		//  7   8      ← leaves at depth 4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}
		root.Right.Right = &TreeNode{Val: 6} // Leaf at depth 3
		root.Left.Left.Left = &TreeNode{Val: 7} // Leaf at depth 4
		root.Left.Right.Left = &TreeNode{Val: 8} // Leaf at depth 4

		result := problem.MinDepth(root)
		expected := 3 // Should find node 6 at depth 3

		if result != expected {
			t.Errorf("Complex unbalanced tree = %d, expected %d", result, expected)
		}
	})
}

func TestMinDepthConstraints(t *testing.T) {
	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test with values at constraint boundaries (values don't affect depth calculation)
		root := &TreeNode{Val: 100000}
		root.Left = &TreeNode{Val: -100000}
		root.Right = &TreeNode{Val: 0}

		result := problem.MinDepth(root)
		expected := 2 // Both left and right are leaves

		if result != expected {
			t.Errorf("Constraint values test = %d, expected %d", result, expected)
		}
	})

	t.Run("Large tree size", func(t *testing.T) {
		// Create a complete binary tree of height 4 (15 nodes)
		input := []*int{
			helper.IntPtr(1),
			helper.IntPtr(2), helper.IntPtr(3),
			helper.IntPtr(4), helper.IntPtr(5), helper.IntPtr(6), helper.IntPtr(7),
			helper.IntPtr(8), helper.IntPtr(9), helper.IntPtr(10), helper.IntPtr(11),
			helper.IntPtr(12), helper.IntPtr(13), helper.IntPtr(14), helper.IntPtr(15),
		}
		root := helper.CreateTreeFromArray(input)

		result := problem.MinDepth(root)
		expected := 4 // All leaves are at level 4

		if result != expected {
			t.Errorf("Large tree test = %d, expected %d", result, expected)
		}
	})
}

func BenchmarkMinDepth(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"MinDepth_Small_7_nodes", 7},
		{"MinDepth_Medium_15_nodes", 15},
		{"MinDepth_Large_31_nodes", 31},
		{"MinDepth_XLarge_63_nodes", 63},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Create a complete binary tree
			values := make([]*int, bm.size)
			for i := 0; i < bm.size; i++ {
				values[i] = helper.IntPtr(i + 1)
			}
			root := helper.CreateTreeFromArray(values)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				problem.MinDepth(root)
			}
		})
	}
}

/***
func ExampleMinDepth() {
	// Create a binary tree:
	//       3
	//      / \
	//     9   20
	//        /  \
	//       15   7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	depth := problem.MinDepth(root)
	println("Minimum depth:", depth)
	// Output: Minimum depth: 2
}
***/

// Test to verify minimum depth property
func TestMinDepthProperties(t *testing.T) {
	t.Run("Minimum depth is always less than or equal to maximum depth", func(t *testing.T) {
		testCases := []struct {
			name string
			tree func() *TreeNode
		}{
			{
				name: "Balanced tree",
				tree: func() *TreeNode {
					return &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 2,
							Left:  &TreeNode{Val: 4},
							Right: &TreeNode{Val: 5},
						},
						Right: &TreeNode{
							Val: 3,
							Left:  &TreeNode{Val: 6},
							Right: &TreeNode{Val: 7},
						},
					}
				},
			},
			{
				name: "Unbalanced tree",
				tree: func() *TreeNode {
					return &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 3,
								Left: &TreeNode{Val: 4},
							},
						},
						Right: &TreeNode{Val: 5},
					}
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				tree := tc.tree()
				minDepth := problem.MinDepth(tree)
				maxDepth := getMaxDepth(tree)

				if minDepth > maxDepth {
					t.Errorf("Min depth (%d) should not be greater than max depth (%d)",
						minDepth, maxDepth)
				}

				if minDepth <= 0 {
					t.Errorf("Min depth should be positive, got %d", minDepth)
				}
			})
		}
	})

	t.Run("Minimum depth equals 1 for single node", func(t *testing.T) {
		root := &TreeNode{Val: 42}
		result := problem.MinDepth(root)

		if result != 1 {
			t.Errorf("Single node should have min depth 1, got %d", result)
		}
	})

	t.Run("Minimum depth equals maximum depth for complete binary tree", func(t *testing.T) {
		// Complete binary tree has all leaves at the same level
		//       1
		//      / \
		//     2   3
		//    / \ / \
		//   4  5 6  7
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}
		root.Right.Left = &TreeNode{Val: 6}
		root.Right.Right = &TreeNode{Val: 7}

		minDepth := problem.MinDepth(root)
		maxDepth := getMaxDepth(root)

		if minDepth != maxDepth {
			t.Errorf("Complete tree should have min depth (%d) equal max depth (%d)",
				minDepth, maxDepth)
		}
	})
}

// Helper function to calculate maximum depth for property testing
func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getMaxDepth(root.Left)
	right := getMaxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}