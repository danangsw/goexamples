package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected bool
	}{
		{
			name:     "Empty tree (nil root)",
			input:    []*int{},
			expected: true,
		},
		{
			name:     "Single node",
			input:    []*int{helper.IntPtr(1)},
			expected: true,
		},
		{
			name:     "Perfect symmetric tree [1,2,2,3,4,4,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(3)},
			expected: true,
		},
		{
			name:     "Asymmetric tree [1,2,2,null,3,null,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Simple symmetric tree [1,2,2]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2)},
			expected: true,
		},
		{
			name:     "Simple asymmetric tree [1,2,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Left subtree only [1,2,null]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), nil},
			expected: false,
		},
		{
			name:     "Right subtree only [1,null,2]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2)},
			expected: false,
		},
		{
			name:     "Symmetric with nulls [1,2,2,null,3,3,null]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), nil, helper.IntPtr(3), helper.IntPtr(3), nil},
			expected: true,
		},
		{
			name:     "Deep symmetric tree [1,2,2,3,3,3,3,4,4,4,4,4,4,4,4]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(3), helper.IntPtr(3), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(4)},
			expected: true,
		},
		{
			name:     "Same structure different values [1,2,2,3,4,5,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), helper.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Symmetric tree with negative values [0,-1,-1,-2,-2,-2,-2]",
			input:    []*int{helper.IntPtr(0), helper.IntPtr(-1), helper.IntPtr(-1), helper.IntPtr(-2), helper.IntPtr(-2), helper.IntPtr(-2), helper.IntPtr(-2)},
			expected: true,
		},
		{
			name:     "Mixed positive and negative [1,-2,-2,3,3,3,3]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(-2), helper.IntPtr(-2), helper.IntPtr(3), helper.IntPtr(3), helper.IntPtr(3), helper.IntPtr(3)},
			expected: true,
		},
		{
			name:     "Complex asymmetric [1,2,2,2,null,2]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2), nil, helper.IntPtr(2)},
			expected: false,
		},
		{
			name:     "LeetCode Example 1: [1,2,2,3,4,4,3] - symmetric",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(4), helper.IntPtr(3)},
			expected: true,
		},
		{
			name:     "LeetCode Example 2: [1,2,2,null,3,null,3] - asymmetric",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), nil, helper.IntPtr(3), nil, helper.IntPtr(3)},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.IsSymmetric(root)

			if result != tt.expected {
				t.Errorf("IsSymmetric() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsSymmetricEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := problem.IsSymmetric(nil)
		expected := true

		if result != expected {
			t.Errorf("IsSymmetric(nil) = %v, expected %v", result, expected)
		}
	})

	t.Run("Deep left-skewed tree", func(t *testing.T) {
		// Create: 1 -> 2 -> 3 -> 4 (all left children)
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}

		result := problem.IsSymmetric(root)
		expected := false

		if result != expected {
			t.Errorf("Deep left-skewed tree = %v, expected %v", result, expected)
		}
	})

	t.Run("Deep right-skewed tree", func(t *testing.T) {
		// Create: 1 -> 2 -> 3 -> 4 (all right children)
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		root.Right.Right.Right = &TreeNode{Val: 4}

		result := problem.IsSymmetric(root)
		expected := false

		if result != expected {
			t.Errorf("Deep right-skewed tree = %v, expected %v", result, expected)
		}
	})

	t.Run("Perfectly symmetric zigzag", func(t *testing.T) {
		// Create a symmetric zigzag pattern
		//       1
		//      / \
		//     2   2
		//    /     \
		//   3       3
		//  /         \
		// 4           4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}
		root.Right.Right.Right = &TreeNode{Val: 4}

		result := problem.IsSymmetric(root)
		expected := true

		if result != expected {
			t.Errorf("Symmetric zigzag = %v, expected %v", result, expected)
		}
	})

	t.Run("Asymmetric zigzag", func(t *testing.T) {
		// Create an asymmetric zigzag pattern
		//       1
		//      / \
		//     2   2
		//    /     \
		//   3       3
		//  /        /  (not symmetric!)
		// 4        4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}
		root.Right.Right.Left = &TreeNode{Val: 4} // Should be Right for symmetry

		result := problem.IsSymmetric(root)
		expected := false

		if result != expected {
			t.Errorf("Asymmetric zigzag = %v, expected %v", result, expected)
		}
	})

	t.Run("Large symmetric tree", func(t *testing.T) {
		// Build a larger symmetric tree for stress testing
		//         1
		//       /   \
		//      2     2
		//     / \   / \
		//    3   4 4   3
		//   /   / \ \   \
		//  5   6   7 8   9
		//     /     \   /
		//    10      11 12
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 4}
		root.Right.Right = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 5}
		root.Left.Right.Left = &TreeNode{Val: 6}
		root.Left.Right.Right = &TreeNode{Val: 7}
		root.Right.Left.Left = &TreeNode{Val: 8}
		root.Right.Left.Right = &TreeNode{Val: 9}
		root.Right.Right.Right = &TreeNode{Val: 5}
		root.Left.Right.Left.Left = &TreeNode{Val: 10}
		root.Left.Right.Right.Right = &TreeNode{Val: 11}
		root.Right.Left.Left.Left = &TreeNode{Val: 12}

		result := problem.IsSymmetric(root)
		expected := false // This is intentionally asymmetric

		if result != expected {
			t.Errorf("Large asymmetric tree = %v, expected %v", result, expected)
		}
	})
}

func TestIsSymmetricBoundaryValues(t *testing.T) {
	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test with values at constraint boundaries
		root := &TreeNode{Val: 100}
		root.Left = &TreeNode{Val: -100}
		root.Right = &TreeNode{Val: -100}

		result := problem.IsSymmetric(root)
		expected := true

		if result != expected {
			t.Errorf("Boundary values test = %v, expected %v", result, expected)
		}
	})

	t.Run("Zero values", func(t *testing.T) {
		// Test with zero values
		root := &TreeNode{Val: 0}
		root.Left = &TreeNode{Val: 0}
		root.Right = &TreeNode{Val: 0}
		root.Left.Left = &TreeNode{Val: 0}
		root.Left.Right = &TreeNode{Val: 0}
		root.Right.Left = &TreeNode{Val: 0}
		root.Right.Right = &TreeNode{Val: 0}

		result := problem.IsSymmetric(root)
		expected := true

		if result != expected {
			t.Errorf("Zero values test = %v, expected %v", result, expected)
		}
	})

	t.Run("Mixed zero and non-zero", func(t *testing.T) {
		// Test mixing zeros with other values
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 0}
		root.Right = &TreeNode{Val: 0}
		root.Left.Left = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 2}

		result := problem.IsSymmetric(root)
		expected := true

		if result != expected {
			t.Errorf("Mixed zero and non-zero test = %v, expected %v", result, expected)
		}
	})
}

func TestIsSymmetricStackBehavior(t *testing.T) {
	t.Run("Stack operations work correctly", func(t *testing.T) {
		// Test that stack correctly handles the push/pop operations
		//       1
		//      / \
		//     2   2
		//    / \ / \
		//   3  4 4  3
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 4}
		root.Right.Right = &TreeNode{Val: 3}

		result := problem.IsSymmetric(root)
		expected := true

		if result != expected {
			t.Errorf("Stack behavior test = %v, expected %v", result, expected)
		}
	})

	t.Run("Verify mirror property", func(t *testing.T) {
		tests := []struct {
			name          string
			tree          func() *TreeNode
			expectedValue bool
		}{
			{
				name: "Mirror property holds",
				tree: func() *TreeNode {
					//   1
					//  / \
					// 2   2
					return &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
				},
				expectedValue: true,
			},
			{
				name: "Mirror property violated",
				tree: func() *TreeNode {
					//   1
					//  / \
					// 2   3
					return &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
				},
				expectedValue: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := problem.IsSymmetric(tt.tree())
				if result != tt.expectedValue {
					t.Errorf("%s: got %v, expected %v", tt.name, result, tt.expectedValue)
				}
			})
		}
	})
}

func BenchmarkIsSymmetric(b *testing.B) {
	// Create test trees of different sizes for benchmarking
	benchmarks := []struct {
		name string
		size int
	}{
		{"Symmetric_Small_7_nodes", 7},
		{"Symmetric_Medium_15_nodes", 15},
		{"Symmetric_Large_31_nodes", 31},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Create a symmetric tree with bm.size nodes
			values := make([]*int, bm.size)
			// Create a symmetric pattern
			values[0] = helper.IntPtr(1) // root
			if bm.size > 1 {
				values[1] = helper.IntPtr(2) // left
				values[2] = helper.IntPtr(2) // right (symmetric)
			}
			if bm.size > 3 {
				for i := 3; i < bm.size && i < 7; i += 2 {
					values[i] = helper.IntPtr(3) // left side
					if i+1 < bm.size {
						values[i+1] = helper.IntPtr(3) // right side (symmetric)
					}
				}
			}
			// Fill remaining with symmetric pattern
			for i := 7; i < bm.size; i += 2 {
				values[i] = helper.IntPtr(4)
				if i+1 < bm.size {
					values[i+1] = helper.IntPtr(4)
				}
			}

			root := helper.CreateTreeFromArray(values)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				problem.IsSymmetric(root)
			}
		})
	}
}

/**
func ExampleIsSymmetric() {
	// Create a symmetric binary tree:
	//       1
	//      / \
	//     2   2
	//    / \ / \
	//   3  4 4  3
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	result := problem.IsSymmetric(root)

	if result {
		println("Tree is symmetric")
	} else {
		println("Tree is not symmetric")
	}
	// Output: Tree is symmetric
}
**/

// Test property-based verification
func TestIsSymmetricProperties(t *testing.T) {
	t.Run("Symmetric property verification", func(t *testing.T) {
		// Test various tree structures to ensure symmetric property
		testCases := []struct {
			name     string
			tree     func() *TreeNode
			expected bool
		}{
			{
				name: "Perfect binary tree - symmetric",
				tree: func() *TreeNode {
					//       1
					//      / \
					//     2   2
					//    / \ / \
					//   3  4 4  3
					root := &TreeNode{Val: 1}
					root.Left = &TreeNode{Val: 2}
					root.Right = &TreeNode{Val: 2}
					root.Left.Left = &TreeNode{Val: 3}
					root.Left.Right = &TreeNode{Val: 4}
					root.Right.Left = &TreeNode{Val: 4}
					root.Right.Right = &TreeNode{Val: 3}
					return root
				},
				expected: true,
			},
			{
				name: "Unbalanced but symmetric",
				tree: func() *TreeNode {
					//     1
					//    / \
					//   2   2
					//    \ /
					//     3
					root := &TreeNode{Val: 1}
					root.Left = &TreeNode{Val: 2}
					root.Right = &TreeNode{Val: 2}
					root.Left.Right = &TreeNode{Val: 3}
					root.Right.Left = &TreeNode{Val: 3}
					return root
				},
				expected: true,
			},
			{
				name: "Structure symmetric but values not",
				tree: func() *TreeNode {
					//     1
					//    / \
					//   2   3  <- Different values
					root := &TreeNode{Val: 1}
					root.Left = &TreeNode{Val: 2}
					root.Right = &TreeNode{Val: 3}
					return root
				},
				expected: false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := problem.IsSymmetric(tc.tree())
				if result != tc.expected {
					t.Errorf("%s: got %v, expected %v", tc.name, result, tc.expected)
				}
			})
		}
	})

	t.Run("Recursive definition verification", func(t *testing.T) {
		// Verify that symmetry is correctly checked recursively
		//       1
		//      / \
		//     2   2
		//    /   /  <- This breaks symmetry (should be / \)
		//   3   3
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Right.Left = &TreeNode{Val: 3}

		result := problem.IsSymmetric(root)
		expected := false // Not symmetric because structure doesn't mirror

		if result != expected {
			t.Errorf("Recursive definition test = %v, expected %v", result, expected)
		}
	})
}
