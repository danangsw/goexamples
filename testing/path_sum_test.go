package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *helper.TreeNode
		targetSum int
		expected  bool
	}{
		{
			name:      "Single node equals target",
			root:      &helper.TreeNode{Val: 5},
			targetSum: 5,
			expected:  true,
		},
		{
			name:      "Single node does not equal target",
			root:      &helper.TreeNode{Val: 5},
			targetSum: 10,
			expected:  false,
		},
		{
			name: "Valid path exists",
			root: &helper.TreeNode{Val: 5,
				Left: &helper.TreeNode{Val: 4,
					Left: &helper.TreeNode{Val: 11,
						Left:  &helper.TreeNode{Val: 7},
						Right: &helper.TreeNode{Val: 2}}},
				Right: &helper.TreeNode{Val: 8,
					Left: &helper.TreeNode{Val: 13},
					Right: &helper.TreeNode{Val: 4,
						Right: &helper.TreeNode{Val: 1}}},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name:      "No valid path",
			root:      &helper.TreeNode{Val: 1, Left: &helper.TreeNode{Val: 2}},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "Empty tree",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
		{
			name: "Negative values in path",
			root: &helper.TreeNode{Val: -2,
				Right: &helper.TreeNode{Val: -3}},
			targetSum: -5,
			expected:  true,
		},
		{
			name: "Multiple paths, only one valid",
			root: &helper.TreeNode{Val: 1,
				Left:  &helper.TreeNode{Val: 2},
				Right: &helper.TreeNode{Val: 3}},
			targetSum: 4,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := problem.HasPathSum(tt.root, tt.targetSum)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
