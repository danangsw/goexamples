package testing_test

import (
	"reflect"
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestPreorderIterativeTraversal(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected []int
	}{
		{
			name:     "Empty tree",
			input:    []*int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			input:    []*int{helper.IntPtr(1)},
			expected: []int{1},
		},
		{
			name:     "Simple tree [1,2,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Tree with left and right children [1,2,3,4,5]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5)},
			expected: []int{1, 2, 4, 5, 3},
		},
		{
			name:     "Left skewed tree [3,2,null,1]",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(2), nil, helper.IntPtr(1)},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Right skewed tree [1,null,2,null,3]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), nil, helper.IntPtr(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complete binary tree [4,2,6,1,3,5,7]",
			input:    []*int{helper.IntPtr(4), helper.IntPtr(2), helper.IntPtr(6), helper.IntPtr(1), helper.IntPtr(3), helper.IntPtr(5), helper.IntPtr(7)},
			expected: []int{4, 2, 1, 3, 6, 5, 7},
		},
		{
			name:     "Complex tree [1,2,3,4,5,null,6,null,null,7,8]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), nil, helper.IntPtr(6), nil, nil, helper.IntPtr(7), helper.IntPtr(8)},
			expected: []int{1, 2, 4, 5, 7, 8, 3, 6},
		},
		{
			name:     "Tree with negative values [0,-1,1,-2,null,null,2]",
			input:    []*int{helper.IntPtr(0), helper.IntPtr(-1), helper.IntPtr(1), helper.IntPtr(-2), nil, nil, helper.IntPtr(2)},
			expected: []int{0, -1, -2, 1, 2},
		},
		{
			name:     "Tree with duplicate values [2,2,2,2]",
			input:    []*int{helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2)},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "Unbalanced tree [5,3,8,2,4,7,9,1]",
			input:    []*int{helper.IntPtr(5), helper.IntPtr(3), helper.IntPtr(8), helper.IntPtr(2), helper.IntPtr(4), helper.IntPtr(7), helper.IntPtr(9), helper.IntPtr(1)},
			expected: []int{5, 3, 2, 1, 4, 8, 7, 9},
		},
		{
			name:     "Two nodes [1,2]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2)},
			expected: []int{1, 2},
		},
		{
			name:     "Two nodes [1,null,2]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2)},
			expected: []int{1, 2},
		},
		{
			name:     "Example from LeetCode: [1,null,2,3]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), helper.IntPtr(3)},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.PreorderIterativeTraversal(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PreorderIterativeTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestPreorderIterativeTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := problem.PreorderIterativeTraversal(nil)
		expected := []int{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PreorderIterativeTraversal(nil) = %v, expected %v", result, expected)
		}
	})

	t.Run("Deep left-heavy tree", func(t *testing.T) {
		// Create a deep left-heavy tree: 5 -> 4 -> 3 -> 2 -> 1
		root := &TreeNode{Val: 5}
		current := root
		for i := 4; i >= 1; i-- {
			current.Left = &TreeNode{Val: i}
			current = current.Left
		}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{5, 4, 3, 2, 1} // Root first, then left chain

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Deep left-heavy tree = %v, expected %v", result, expected)
		}
	})

	t.Run("Deep right-heavy tree", func(t *testing.T) {
		// Create a deep right-heavy tree: 1 -> 2 -> 3 -> 4 -> 5
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 5; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
		}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4, 5} // Root first, then right chain

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Deep right-heavy tree = %v, expected %v", result, expected)
		}
	})

	t.Run("Alternating left-right structure", func(t *testing.T) {
		// Create: 1 -> left: 2 -> right: 3 -> left: 4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 3}
		root.Left.Right.Left = &TreeNode{Val: 4}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4} // Root first pattern

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Alternating structure = %v, expected %v", result, expected)
		}
	})

	t.Run("Single left child", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{1, 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Single left child = %v, expected %v", result, expected)
		}
	})

	t.Run("Single right child", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{1, 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Single right child = %v, expected %v", result, expected)
		}
	})
}

func TestPreorderIterativeTraversalLargeValues(t *testing.T) {
	t.Run("Large values within constraints", func(t *testing.T) {
		// Test with values at the constraint boundaries [-100, 100]
		input := []*int{helper.IntPtr(0), helper.IntPtr(-100), helper.IntPtr(100), helper.IntPtr(-50), helper.IntPtr(50)}
		root := helper.CreateTreeFromArray(input)

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{0, -100, -50, 50, 100} // Preorder: root first

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Large values test = %v, expected %v", result, expected)
		}
	})

	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test extreme values
		root := &TreeNode{Val: 100}
		root.Left = &TreeNode{Val: -100}
		root.Right = &TreeNode{Val: 0}

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{100, -100, 0}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Maximum constraint values = %v, expected %v", result, expected)
		}
	})
}

func TestPreorderIterativeTraversalStackBehavior(t *testing.T) {
	t.Run("Stack operations work correctly", func(t *testing.T) {
		// Create a tree that will test stack push/pop operations thoroughly
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

		result := problem.PreorderIterativeTraversal(root)
		expected := []int{1, 2, 4, 5, 3, 6, 7} // Preorder traversal

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Stack behavior test = %v, expected %v", result, expected)
		}
	})

	t.Run("Verify root-first property", func(t *testing.T) {
		// Test that root is always processed first in each subtree
		tests := []struct {
			name          string
			tree          func() *TreeNode
			expectedFirst int
		}{
			{
				name: "Simple tree",
				tree: func() *TreeNode {
					return &TreeNode{Val: 42, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
				},
				expectedFirst: 42,
			},
			{
				name: "Large root value",
				tree: func() *TreeNode {
					return &TreeNode{Val: 100, Left: &TreeNode{Val: -100}}
				},
				expectedFirst: 100,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := problem.PreorderIterativeTraversal(tt.tree())
				if len(result) == 0 || result[0] != tt.expectedFirst {
					t.Errorf("First element should be %d, got %v", tt.expectedFirst, result)
				}
			})
		}
	})
}

func BenchmarkPreorderIterativeTraversal(b *testing.B) {
	// Create test trees of different sizes
	benchmarks := []struct {
		name string
		size int
	}{
		{"Preorder_Small_7_nodes", 7},
		{"Preorder_Medium_31_nodes", 31},
		{"Preorder_Large_100_nodes", 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Create a complete binary tree with bm.size nodes
			values := make([]*int, bm.size)
			for i := 0; i < bm.size; i++ {
				values[i] = helper.IntPtr(i + 1)
			}
			root := helper.CreateTreeFromArray(values)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				problem.PreorderIterativeTraversal(root)
			}
		})
	}
}

/****
func ExamplePreorderIterativeTraversal() {
	// Create a simple binary tree:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	result := problem.PreorderIterativeTraversal(root)

	for i, val := range result {
		if i > 0 {
			print(", ")
		}
		print(val)
	}
	println()
	// Output: 1, 2, 4, 5, 3
}
***/

// Test to verify preorder property: root processed before children
func TestPreorderIterativeTraversalProperties(t *testing.T) {
	t.Run("Root processed before all children", func(t *testing.T) {
		// Create tree where we can verify parent-child ordering
		root := &TreeNode{Val: 10}
		root.Left = &TreeNode{Val: 5}
		root.Right = &TreeNode{Val: 15}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 7}

		result := problem.PreorderIterativeTraversal(root)

		// Verify root (10) comes before all its descendants
		rootIndex := -1
		for i, val := range result {
			if val == 10 {
				rootIndex = i
				break
			}
		}

		if rootIndex != 0 {
			t.Errorf("Root should be first element, found at index %d", rootIndex)
		}

		// Verify each parent comes before its children
		expectedOrder := []int{10, 5, 3, 7, 15}
		if !reflect.DeepEqual(result, expectedOrder) {
			t.Errorf("Preorder property violated: got %v, expected %v", result, expectedOrder)
		}
	})

	t.Run("Compare with manual preorder calculation", func(t *testing.T) {
		// Test multiple tree structures to verify correctness
		testCases := []struct {
			name     string
			tree     func() *TreeNode
			expected []int
		}{
			{
				name: "Binary search tree",
				tree: func() *TreeNode {
					//       8
					//      / \
					//     3   10
					//    / \    \
					//   1   6    14
					//      / \   /
					//     4   7 13
					root := &TreeNode{Val: 8}
					root.Left = &TreeNode{Val: 3}
					root.Right = &TreeNode{Val: 10}
					root.Left.Left = &TreeNode{Val: 1}
					root.Left.Right = &TreeNode{Val: 6}
					root.Right.Right = &TreeNode{Val: 14}
					root.Left.Right.Left = &TreeNode{Val: 4}
					root.Left.Right.Right = &TreeNode{Val: 7}
					root.Right.Right.Left = &TreeNode{Val: 13}
					return root
				},
				expected: []int{8, 3, 1, 6, 4, 7, 10, 14, 13},
			},
			{
				name: "Asymmetric tree",
				tree: func() *TreeNode {
					//     1
					//    / \
					//   2   3
					//  /     \
					// 4       5
					//          \
					//           6
					root := &TreeNode{Val: 1}
					root.Left = &TreeNode{Val: 2}
					root.Right = &TreeNode{Val: 3}
					root.Left.Left = &TreeNode{Val: 4}
					root.Right.Right = &TreeNode{Val: 5}
					root.Right.Right.Right = &TreeNode{Val: 6}
					return root
				},
				expected: []int{1, 2, 4, 3, 5, 6},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := problem.PreorderIterativeTraversal(tc.tree())
				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("%s: got %v, expected %v", tc.name, result, tc.expected)
				}
			})
		}
	})

	t.Run("Verify preorder vs inorder difference", func(t *testing.T) {
		// Same tree should produce different results for preorder vs inorder
		//       4
		//      / \
		//     2   6
		//    / \ / \
		//   1  3 5  7
		root := &TreeNode{Val: 4}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 6}
		root.Left.Left = &TreeNode{Val: 1}
		root.Left.Right = &TreeNode{Val: 3}
		root.Right.Left = &TreeNode{Val: 5}
		root.Right.Right = &TreeNode{Val: 7}

		preorderResult := problem.PreorderIterativeTraversal(root)
		inorderResult := problem.InorderIterativeTraversal(root)

		expectedPreorder := []int{4, 2, 1, 3, 6, 5, 7}
		expectedInorder := []int{1, 2, 3, 4, 5, 6, 7}

		if !reflect.DeepEqual(preorderResult, expectedPreorder) {
			t.Errorf("Preorder result = %v, expected %v", preorderResult, expectedPreorder)
		}

		if !reflect.DeepEqual(inorderResult, expectedInorder) {
			t.Errorf("Inorder result = %v, expected %v", inorderResult, expectedInorder)
		}

		// They should be different for this tree
		if reflect.DeepEqual(preorderResult, inorderResult) {
			t.Error("Preorder and inorder results should be different for this tree structure")
		}
	})
}
