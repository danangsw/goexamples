package testing_test

import (
	"reflect"
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestInorderIterativeTraversal(t *testing.T) {
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
			name:     "Example 1: [1,null,2,3]",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), helper.IntPtr(3)},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2: [1,2,3,4,5,null,8,null,null,6,7,9]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(4), helper.IntPtr(5), nil, helper.IntPtr(8), nil, nil, helper.IntPtr(6), helper.IntPtr(7), helper.IntPtr(9)},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name:     "Left skewed tree",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(2), nil, helper.IntPtr(1)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Right skewed tree",
			input:    []*int{helper.IntPtr(1), nil, helper.IntPtr(2), nil, helper.IntPtr(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complete binary tree",
			input:    []*int{helper.IntPtr(4), helper.IntPtr(2), helper.IntPtr(6), helper.IntPtr(1), helper.IntPtr(3), helper.IntPtr(5), helper.IntPtr(7)},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Tree with negative values",
			input:    []*int{helper.IntPtr(0), helper.IntPtr(-1), helper.IntPtr(1), helper.IntPtr(-2), nil, nil, helper.IntPtr(2)},
			expected: []int{-2, -1, 0, 1, 2},
		},
		{
			name:     "Tree with duplicate values",
			input:    []*int{helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(2)},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "Complex unbalanced tree",
			input:    []*int{helper.IntPtr(5), helper.IntPtr(3), helper.IntPtr(8), helper.IntPtr(2), helper.IntPtr(4), helper.IntPtr(7), helper.IntPtr(9), helper.IntPtr(1)},
			expected: []int{1, 2, 3, 4, 5, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.InorderIterativeTraversal(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InorderIterativeTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestInorderIterativeTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := problem.InorderIterativeTraversal(nil)
		expected := []int{}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InorderIterativeTraversal(nil) = %v, expected %v", result, expected)
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

		result := problem.InorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4, 5}

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

		result := problem.InorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Deep right-heavy tree = %v, expected %v", result, expected)
		}
	})
}

func TestInorderIterativeTraversalLargeValues(t *testing.T) {
	t.Run("Large values within constraints", func(t *testing.T) {
		// Test with values at the constraint boundaries [-100, 100]
		input := []*int{helper.IntPtr(0), helper.IntPtr(-100), helper.IntPtr(100), helper.IntPtr(-50), helper.IntPtr(50)}
		root := helper.CreateTreeFromArray(input)

		result := problem.InorderIterativeTraversal(root)
		expected := []int{-50, -100, 50, 0, 100}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Large values test = %v, expected %v", result, expected)
		}
	})
}

func BenchmarkInorderIterativeTraversal(b *testing.B) {
	// Create test trees of different sizes
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small_7_nodes", 7},
		{"Medium_31_nodes", 31},
		{"Large_100_nodes", 100},
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
				problem.InorderIterativeTraversal(root)
			}
		})
	}
}

// Helper test to verify stack behavior
func TestInorderIterativeTraversalStackBehavior(t *testing.T) {
	t.Run("Stack operations work correctly", func(t *testing.T) {
		// Create a tree that will test stack push/pop operations thoroughly
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

		result := problem.InorderIterativeTraversal(root)
		expected := []int{1, 2, 3, 4, 5, 6, 7}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Stack behavior test = %v, expected %v", result, expected)
		}
	})
}

/**
func ExampleInorderIterativeTraversal() {
	// Create a simple binary tree:
	//       2
	//      / \
	//     1   3
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	result := InorderIterativeTraversal(root)

	for i, val := range result {
		if i > 0 {
			print(", ")
		}
		print(val)
	}
	println()
	// Output: 1, 2, 3
}
	**/
