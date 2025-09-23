package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestIsBalanced(t *testing.T) {
	// Balanced tree:
	//       3
	//      / \
	//     2   4
	//    /     \
	//   1       5
	balanced := &helper.TreeNode{
		Val: 3,
		Left: &helper.TreeNode{
			Val:   2,
			Left:  &helper.TreeNode{Val: 1, Left: nil, Right: nil},
			Right: nil,
		},
		Right: &helper.TreeNode{
			Val:   4,
			Left:  nil,
			Right: &helper.TreeNode{Val: 5, Left: nil, Right: nil},
		},
	}

	// Unbalanced tree:
	//   1
	//  /
	// 2
	///
	//3
	unbalanced := &helper.TreeNode{
		Val: 1,
		Left: &helper.TreeNode{
			Val: 2,
			Left: &helper.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}

	tests := []struct {
		name     string
		root     *helper.TreeNode
		expected bool
	}{
		{
			name:     "Balanced tree",
			root:     balanced,
			expected: true,
		},
		{
			name:     "Unbalanced tree",
			root:     unbalanced,
			expected: false,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     &helper.TreeNode{Val: 1, Left: nil, Right: nil},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.IsBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("IsBalanced() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsBalancedEdgeCases(t *testing.T) {
	t.Run("Perfect binary tree", func(t *testing.T) {
		// Perfect binary tree (height 3):
		//       1
		//      / \
		//     2   3
		//    / \ / \
		//   4  5 6  7
		perfect := &helper.TreeNode{
			Val: 1,
			Left: &helper.TreeNode{
				Val:   2,
				Left:  &helper.TreeNode{Val: 4, Left: nil, Right: nil},
				Right: &helper.TreeNode{Val: 5, Left: nil, Right: nil},
			},
			Right: &helper.TreeNode{
				Val:   3,
				Left:  &helper.TreeNode{Val: 6, Left: nil, Right: nil},
				Right: &helper.TreeNode{Val: 7, Left: nil, Right: nil},
			},
		}

		if !problem.IsBalanced(perfect) {
			t.Errorf("Perfect binary tree should be balanced")
		}
	})

	t.Run("Right-heavy unbalanced tree", func(t *testing.T) {
		// Right-heavy unbalanced tree:
		//   1
		//    \
		//     2
		//      \
		//       3
		//        \
		//         4
		rightHeavy := &helper.TreeNode{
			Val:  1,
			Left: nil,
			Right: &helper.TreeNode{
				Val:  2,
				Left: nil,
				Right: &helper.TreeNode{
					Val:  3,
					Left: nil,
					Right: &helper.TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
		}

		if problem.IsBalanced(rightHeavy) {
			t.Errorf("Right-heavy tree should be unbalanced")
		}
	})

	t.Run("Balanced but not perfect", func(t *testing.T) {
		// Balanced but not perfect:
		//     1
		//    / \
		//   2   3
		//  /
		// 4
		balancedNotPerfect := &helper.TreeNode{
			Val: 1,
			Left: &helper.TreeNode{
				Val:   2,
				Left:  &helper.TreeNode{Val: 4, Left: nil, Right: nil},
				Right: nil,
			},
			Right: &helper.TreeNode{Val: 3, Left: nil, Right: nil},
		}

		if !problem.IsBalanced(balancedNotPerfect) {
			t.Errorf("Tree with height difference of 1 should be balanced")
		}
	})

	t.Run("Complex balanced tree", func(t *testing.T) {
		// More complex balanced tree:
		//        1
		//       / \
		//      2   3
		//     / \   \
		//    4   5   6
		//   /   / \
		//  7   8   9
		complex := &helper.TreeNode{
			Val: 1,
			Left: &helper.TreeNode{
				Val: 2,
				Left: &helper.TreeNode{
					Val:   4,
					Left:  &helper.TreeNode{Val: 7, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &helper.TreeNode{
					Val:   5,
					Left:  &helper.TreeNode{Val: 8, Left: nil, Right: nil},
					Right: &helper.TreeNode{Val: 9, Left: nil, Right: nil},
				},
			},
			Right: &helper.TreeNode{
				Val:   3,
				Left:  nil,
				Right: &helper.TreeNode{Val: 6, Left: nil, Right: nil},
			},
		}

		if !problem.IsBalanced(complex) {
			t.Errorf("Complex balanced tree should be balanced")
		}
	})
}

func TestIsBalancedWithArrayInput(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected bool
	}{
		{
			name:     "Balanced tree from array [3,9,20,null,null,15,7]",
			input:    []*int{helper.IntPtr(3), helper.IntPtr(9), helper.IntPtr(20), nil, nil, helper.IntPtr(15), helper.IntPtr(7)},
			expected: true,
		},
		{
			name:     "Unbalanced tree from array [1,2,2,3,3,null,null,4,4]",
			input:    []*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(2), helper.IntPtr(3), helper.IntPtr(3), nil, nil, helper.IntPtr(4), helper.IntPtr(4)},
			expected: false,
		},
		{
			name:     "Single node from array [1]",
			input:    []*int{helper.IntPtr(1)},
			expected: true,
		},
		{
			name:     "Empty array",
			input:    []*int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := helper.CreateTreeFromArray(tt.input)
			result := problem.IsBalanced(root)
			if result != tt.expected {
				t.Errorf("IsBalanced() = %v, expected %v for input %v", result, tt.expected, tt.input)
			}
		})
	}
}
