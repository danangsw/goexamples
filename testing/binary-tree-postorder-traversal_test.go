package testing_test

import (
	"reflect"
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestPostorderIterativeTraversal(t *testing.T) {
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
			expected: []int{2, 3, 1}, // Left, Right, Root
		},
		{
			name:     "Tree with left and right children [1,2,3,4,5]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5)},
			expected: []int{4, 5, 2, 3, 1}, // Postorder: children before parent
		},
		{
			name:     "Left skewed tree [3,2,null,1]",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(2), nil, helper.IntPtr(1)},
			expected: []int{1, 2, 3}, // Bottom-up processing
		},
		{
			name:     "Right skewed tree [1,null,2,null,3]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), nil, helper.IntPtr(3)},
			expected: []int{3, 2, 1}, // Bottom-up processing
		},
		{
			name:     "Complete binary tree [4,2,6,1,3,5,7]",
			input:    []*int{helper.IntPtr(4), helper.IntPtr(2), helper.IntPtr(6), helper.IntPtr(1), helper.IntPtr(3), helper.IntPtr(5), helper.IntPtr(7)},
			expected: []int{1, 3, 2, 5, 7, 6, 4}, // Leaves first, then parents
		},
		{
			name:     "Complex tree [1,2,3,4,5,null,6,null,null,7,8]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), nil, helper.IntPtr(6), nil, nil, helper.IntPtr(7), helper.IntPtr(8)},
			expected: []int{4, 7, 8, 5, 2, 6, 3, 1}, // Deep leaves first
		},
		{
			name:     "Tree with negative values [0,-1,1,-2,null,null,2]",
			input:    []*int{helper.IntPtr(0), helper.IntPtr(-1), helper.IntPtr(1), helper.IntPtr(-2), nil, nil, helper.IntPtr(2)},
			expected: []int{-2, -1, 2, 1, 0}, // Postorder with negative values
		},
		{
			name:     "Tree with duplicate values [2,2,2,2]",
			input:    []*int{helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2)},
			expected: []int{2, 2, 2, 2}, // All nodes processed bottom-up
		},
		{
			name:     "Unbalanced tree [5,3,8,2,4,7,9,1]",
			input:    []*int{helper.IntPtr(5), helper.IntPtr(3), helper.IntPtr(8), helper.IntPtr(2), helper.IntPtr(4), helper.IntPtr(7), helper.IntPtr(9), helper.IntPtr(1)},
			expected: []int{1, 2, 4, 3, 7, 9, 8, 5}, // Complex postorder
		},
		{
			name:     "Two nodes [1,2]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2)},
			expected: []int{2, 1}, // Child before parent
		},
		{
			name:     "Two nodes [1,null,2]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2)},
			expected: []int{2, 1}, // Right child before parent
		},
		{
			name:     "Example from LeetCode: [1,null,2,3]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), helper.IntPtr(3)},
			expected: []int{3, 2, 1}, // Postorder: 3, 2, 1
		},
		{
			name:     "Deep left subtree [1,2,null,3,null,4]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4)},
			expected: []int{4, 3, 2, 1}, // Deep left chain
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.PostorderIterativeTraversal(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PostorderIterativeTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestPostorderIterativeTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := problem.PostorderIterativeTraversal(nil)
		expected := []int{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PostorderIterativeTraversal(nil) = %v, expected %v", result, expected)
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

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4, 5} // Bottom-up: deepest first

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

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{5, 4, 3, 2, 1} // Bottom-up: deepest first

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

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{4, 3, 2, 1} // Bottom-up traversal

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Alternating structure = %v, expected %v", result, expected)
		}
	})

	t.Run("Single left child", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{2, 1} // Child before parent

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Single left child = %v, expected %v", result, expected)
		}
	})

	t.Run("Single right child", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{2, 1} // Child before parent

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Single right child = %v, expected %v", result, expected)
		}
	})

	t.Run("Only left subtree", func(t *testing.T) {
		//     1
		//    /
		//   2
		//  / \
		// 3   4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 4}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{3, 4, 2, 1} // Process left subtree completely first

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Only left subtree = %v, expected %v", result, expected)
		}
	})

	t.Run("Only right subtree", func(t *testing.T) {
		// 1
		//  \
		//   2
		//  / \
		// 3   4
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 4}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{3, 4, 2, 1} // Process right subtree completely first

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Only right subtree = %v, expected %v", result, expected)
		}
	})
}

func TestPostorderIterativeTraversalLargeValues(t *testing.T) {
	t.Run("Large values within constraints", func(t *testing.T) {
		// Test with values at the constraint boundaries [-100, 100]
		input := []*int{helper.IntPtr(0), helper.IntPtr(-100), helper.IntPtr(100), helper.IntPtr(-50), helper.IntPtr(50)}
		root := helper.CreateTreeFromArray(input)

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{-50, 50, -100, 100, 0} // Postorder: children before parent

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Large values test = %v, expected %v", result, expected)
		}
	})

	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test extreme values
		root := &TreeNode{Val: 100}
		root.Left = &TreeNode{Val: -100}
		root.Right = &TreeNode{Val: 0}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{-100, 0, 100} // Children before parent

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Maximum constraint values = %v, expected %v", result, expected)
		}
	})
}

func TestPostorderIterativeTraversalStackBehavior(t *testing.T) {
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

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{4, 5, 2, 6, 7, 3, 1} // Postorder: leaves to root

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Stack behavior test = %v, expected %v", result, expected)
		}
	})

	t.Run("Verify children-before-parent property", func(t *testing.T) {
		// Test that children are always processed before their parent
		tests := []struct {
			name         string
			tree         func() *TreeNode
			expectedLast int
		}{
			{
				name: "Simple tree",
				tree: func() *TreeNode {
					return &TreeNode{Val: 42, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
				},
				expectedLast: 42,
			},
			{
				name: "Large root value",
				tree: func() *TreeNode {
					return &TreeNode{Val: 100, Left: &TreeNode{Val: -100}}
				},
				expectedLast: 100,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := problem.PostorderIterativeTraversal(tt.tree())
				if len(result) == 0 || result[len(result)-1] != tt.expectedLast {
					t.Errorf("Last element should be %d, got %v", tt.expectedLast, result)
				}
			})
		}
	})
}

func BenchmarkPostorderIterativeTraversal(b *testing.B) {
	// Create test trees of different sizes
	benchmarks := []struct {
		name string
		size int
	}{
		{"Postorder_Small_7_nodes", 7},
		{"Postorder_Medium_31_nodes", 31},
		{"Postorder_Large_100_nodes", 100},
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
				problem.PostorderIterativeTraversal(root)
			}
		})
	}
}

/***
func ExamplePostorderIterativeTraversal() {
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

	result := problem.PostorderIterativeTraversal(root)

	for i, val := range result {
		if i > 0 {
			print(", ")
		}
		print(val)
	}
	println()
	// Output: 4, 5, 2, 3, 1
}
***/

// Test to verify postorder property: children processed before parent
func TestPostorderIterativeTraversalProperties(t *testing.T) {
	t.Run("Children processed before parent", func(t *testing.T) {
		// Create tree where we can verify child-parent ordering
		root := &TreeNode{Val: 10}
		root.Left = &TreeNode{Val: 5}
		root.Right = &TreeNode{Val: 15}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 7}

		result := problem.PostorderIterativeTraversal(root)

		// Verify root (10) comes after all its descendants
		rootIndex := -1
		for i, val := range result {
			if val == 10 {
				rootIndex = i
				break
			}
		}

		if rootIndex != len(result)-1 {
			t.Errorf("Root should be last element, found at index %d", rootIndex)
		}

		// Verify each parent comes after its children
		expectedOrder := []int{3, 7, 5, 15, 10}
		if !reflect.DeepEqual(result, expectedOrder) {
			t.Errorf("Postorder property violated: got %v, expected %v", result, expectedOrder)
		}
	})

	t.Run("Compare with manual postorder calculation", func(t *testing.T) {
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
				expected: []int{1, 4, 7, 6, 3, 13, 14, 10, 8},
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
				expected: []int{4, 2, 6, 5, 3, 1},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := problem.PostorderIterativeTraversal(tc.tree())
				if !reflect.DeepEqual(result, tc.expected) {
					t.Errorf("%s: got %v, expected %v", tc.name, result, tc.expected)
				}
			})
		}
	})

	t.Run("Compare all three traversals", func(t *testing.T) {
		// Same tree should produce different results for all three traversals
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
		postorderResult := problem.PostorderIterativeTraversal(root)

		expectedPreorder := []int{4, 2, 1, 3, 6, 5, 7}  // Root first
		expectedInorder := []int{1, 2, 3, 4, 5, 6, 7}   // Sorted for BST
		expectedPostorder := []int{1, 3, 2, 5, 7, 6, 4} // Children first

		if !reflect.DeepEqual(preorderResult, expectedPreorder) {
			t.Errorf("Preorder result = %v, expected %v", preorderResult, expectedPreorder)
		}

		if !reflect.DeepEqual(inorderResult, expectedInorder) {
			t.Errorf("Inorder result = %v, expected %v", inorderResult, expectedInorder)
		}

		if !reflect.DeepEqual(postorderResult, expectedPostorder) {
			t.Errorf("Postorder result = %v, expected %v", postorderResult, expectedPostorder)
		}

		// All three should be different for this tree
		if reflect.DeepEqual(preorderResult, inorderResult) {
			t.Error("Preorder and inorder results should be different")
		}
		if reflect.DeepEqual(preorderResult, postorderResult) {
			t.Error("Preorder and postorder results should be different")
		}
		if reflect.DeepEqual(inorderResult, postorderResult) {
			t.Error("Inorder and postorder results should be different")
		}
	})

	t.Run("Verify leaf-first property", func(t *testing.T) {
		// In postorder, all leaves should appear before their parents
		//       1
		//      / \
		//     2   3
		//    /   / \
		//   4   5   6
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 5}
		root.Right.Right = &TreeNode{Val: 6}

		result := problem.PostorderIterativeTraversal(root)
		expected := []int{4, 2, 5, 6, 3, 1}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Leaf-first property test = %v, expected %v", result, expected)
		}

		// Verify that leaves (4, 5, 6) appear before their parents
		leafPositions := make(map[int]int)
		for i, val := range result {
			if val == 4 || val == 5 || val == 6 {
				leafPositions[val] = i
			}
		}

		// Parent positions
		parentPositions := make(map[int]int)
		for i, val := range result {
			if val == 1 || val == 2 || val == 3 {
				parentPositions[val] = i
			}
		}

		// Check that leaf 4 comes before its parent 2
		if leafPositions[4] >= parentPositions[2] {
			t.Error("Leaf 4 should come before parent 2")
		}

		// Check that leaves 5,6 come before their parent 3
		if leafPositions[5] >= parentPositions[3] || leafPositions[6] >= parentPositions[3] {
			t.Error("Leaves 5,6 should come before parent 3")
		}

		// Check that all internal nodes come before root 1
		if parentPositions[2] >= parentPositions[1] || parentPositions[3] >= parentPositions[1] {
			t.Error("All internal nodes should come before root 1")
		}
	})
}
