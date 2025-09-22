package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestMaxDepthRecursive(t *testing.T) {
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
			name:     "LeetCode Example 1: [3,9,20,null,null,15,7]",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(9), helper.IntPtr(20), nil, nil, helper.IntPtr(15), helper.IntPtr(7)},
			expected: 3,
		},
		{
			name:     "LeetCode Example 2: [1,null,2]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2)},
			expected: 2,
		},
		{
			name:     "Left skewed tree [1,2,null,3,null,4]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4)},
			expected: 4,
		},
		{
			name:     "Right skewed tree [1,null,2,null,3,null,4]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(4)},
			expected: 4,
		},
		{
			name:     "Complete binary tree [1,2,3,4,5,6,7]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), helper.IntPtr(6), helper.IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Unbalanced tree [1,2,3,4,null,null,7,8]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), nil, nil, helper.IntPtr(7), helper.IntPtr(8)},
			expected: 4,
		},
		{
			name:     "Deep left subtree [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), helper.IntPtr(6), helper.IntPtr(7), helper.IntPtr(8), helper.IntPtr(9), helper.IntPtr(10), helper.IntPtr(11), helper.IntPtr(12), helper.IntPtr(13), helper.IntPtr(14), helper.IntPtr(15)},
			expected: 4,
		},
		{
			name:     "Two levels [1,2,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3)},
			expected: 2,
		},
		{
			name:     "Asymmetric tree [1,2,3,4,5,null,6,7,null,null,8]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), nil, helper.IntPtr(6), helper.IntPtr(7), nil, nil, helper.IntPtr(8)},
			expected: 4,
		},
	}

	// Test all implementations with the same test cases
	implementations := map[string]func(*TreeNode) int{
		"Recursive":             problem.MaxDepthRecursive,
		"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
		"IterativeBFS":          problem.MaxDepthIterativeBFS,
		"IterativeDFS":          problem.MaxDepthIterativeDFS,
		"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
	}

	for implName, impl := range implementations {
		t.Run(implName, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					root := helper.CreateTreeFromArray(tt.input)
					result := impl(root)

					if result != tt.expected {
						t.Errorf("%s() = %d, expected %d", implName, result, tt.expected)
					}
				})
			}
		})
	}
}

func TestMaxDepthEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(nil)
			expected := 0

			if result != expected {
				t.Errorf("%s(nil) = %d, expected %d", implName, result, expected)
			}
		}
	})

	t.Run("Very deep left-heavy tree", func(t *testing.T) {
		// Create: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Left = &TreeNode{Val: i}
			current = current.Left
		}

		expected := 10
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(root)

			if result != expected {
				t.Errorf("%s() = %d, expected %d", implName, result, expected)
			}
		}
	})

	t.Run("Very deep right-heavy tree", func(t *testing.T) {
		// Create: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
		}

		expected := 10
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(root)

			if result != expected {
				t.Errorf("%s() = %d, expected %d", implName, result, expected)
			}
		}
	})

	t.Run("Perfectly balanced tree", func(t *testing.T) {
		// Create a perfectly balanced tree:
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

		expected := 3
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(root)

			if result != expected {
				t.Errorf("%s() = %d, expected %d", implName, result, expected)
			}
		}
	})

	t.Run("Zigzag pattern tree", func(t *testing.T) {
		// Create a zigzag pattern:
		//     1
		//    /
		//   2
		//    \
		//     3
		//    /
		//   4
		//    \
		//     5
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 3}
		root.Left.Right.Left = &TreeNode{Val: 4}
		root.Left.Right.Left.Right = &TreeNode{Val: 5}

		expected := 5
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(root)

			if result != expected {
				t.Errorf("%s() = %d, expected %d", implName, result, expected)
			}
		}
	})
}

func TestMaxDepthConstraints(t *testing.T) {
	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test with values at constraint boundaries (typically node values don't affect depth)
		root := &TreeNode{Val: 100}
		root.Left = &TreeNode{Val: -100}
		root.Right = &TreeNode{Val: 0}
		root.Left.Left = &TreeNode{Val: 50}

		expected := 3
		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for implName, impl := range implementations {
			result := impl(root)

			if result != expected {
				t.Errorf("%s() = %d, expected %d", implName, result, expected)
			}
		}
	})
}

func BenchmarkMaxDepth(b *testing.B) {
	// Create test trees of different sizes
	benchmarks := []struct {
		name string
		size int
	}{
		{"MaxDepth_Small_7_nodes", 7},
		{"MaxDepth_Medium_15_nodes", 15},
		{"MaxDepth_Large_31_nodes", 31},
		{"MaxDepth_XLarge_63_nodes", 63},
	}

	implementations := map[string]func(*TreeNode) int{
		"Recursive":             problem.MaxDepthRecursive,
		"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
		"IterativeBFS":          problem.MaxDepthIterativeBFS,
		"IterativeDFS":          problem.MaxDepthIterativeDFS,
		"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
	}

	for _, bm := range benchmarks {
		// Create a complete binary tree
		values := make([]*int, bm.size)
		for i := 0; i < bm.size; i++ {
			values[i] = helper.IntPtr(i + 1)
		}
		root := helper.CreateTreeFromArray(values)

		for implName, impl := range implementations {
			b.Run(bm.name+"_"+implName, func(b *testing.B) {
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					impl(root)
				}
			})
		}
	}
}

/**
func ExampleMaxDepthRecursive() {
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

	depth := problem.MaxDepthRecursive(root)
	println("Maximum depth:", depth)
	// Output: Maximum depth: 3
}
**/

// Test to verify all implementations produce identical results
func TestMaxDepthConsistency(t *testing.T) {
	t.Run("All implementations should produce identical results", func(t *testing.T) {
		testCases := []struct {
			name string
			tree func() *TreeNode
		}{
			{
				name: "Complex asymmetric tree",
				tree: func() *TreeNode {
					//         1
					//       /   \
					//      2     3
					//     / \     \
					//    4   5     6
					//   /   / \   /
					//  7   8   9 10
					//     /
					//    11
					root := &TreeNode{Val: 1}
					root.Left = &TreeNode{Val: 2}
					root.Right = &TreeNode{Val: 3}
					root.Left.Left = &TreeNode{Val: 4}
					root.Left.Right = &TreeNode{Val: 5}
					root.Right.Right = &TreeNode{Val: 6}
					root.Left.Left.Left = &TreeNode{Val: 7}
					root.Left.Right.Left = &TreeNode{Val: 8}
					root.Left.Right.Right = &TreeNode{Val: 9}
					root.Right.Right.Left = &TreeNode{Val: 10}
					root.Left.Right.Left.Left = &TreeNode{Val: 11}
					return root
				},
			},
			{
				name: "Single branch tree",
				tree: func() *TreeNode {
					root := &TreeNode{Val: 1}
					root.Right = &TreeNode{Val: 2}
					root.Right.Left = &TreeNode{Val: 3}
					root.Right.Left.Right = &TreeNode{Val: 4}
					return root
				},
			},
		}

		implementations := map[string]func(*TreeNode) int{
			"Recursive":             problem.MaxDepthRecursive,
			"RecursiveOneLiner":     problem.MaxDepthRecursiveOneLiner,
			"IterativeBFS":          problem.MaxDepthIterativeBFS,
			"IterativeDFS":          problem.MaxDepthIterativeDFS,
			"IterativeDFSOptimized": problem.MaxDepthIterativeDFSOptimized,
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				tree := tc.tree()
				results := make(map[string]int)

				// Get results from all implementations
				for implName, impl := range implementations {
					results[implName] = impl(tree)
				}

				// Verify all results are identical
				var expected int
				var firstImpl string
				for implName, result := range results {
					if firstImpl == "" {
						firstImpl = implName
						expected = result
					} else if result != expected {
						t.Errorf("Implementation %s returned %d, but %s returned %d",
							implName, result, firstImpl, expected)
					}
				}
			})
		}
	})
}
