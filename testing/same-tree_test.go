package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		tree1    []*int
		tree2    []*int
		expected bool
	}{
		{"Both nil", []*int{}, []*int{}, true},
		{"One nil", []*int{helper.IntPtr(1)}, []*int{}, false},
		{"Same single node", []*int{helper.IntPtr(1)}, []*int{helper.IntPtr(1)}, true},
		{"Different single node", []*int{helper.IntPtr(1)}, []*int{helper.IntPtr(2)}, false},
		{"Same structure and values",
			[]*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3)},
			[]*int{helper.IntPtr(1), helper.IntPtr(2), helper.IntPtr(3)}, true},
		{"Same structure, different values",
			[]*int{helper.IntPtr(1), helper.IntPtr(2)},
			[]*int{helper.IntPtr(1), helper.IntPtr(3)}, false},
		{"Different structure",
			[]*int{helper.IntPtr(1), helper.IntPtr(2)},
			[]*int{helper.IntPtr(1), nil, helper.IntPtr(2)}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree1 := helper.CreateTreeFromArray(tt.tree1)
			tree2 := helper.CreateTreeFromArray(tt.tree2)

			result := problem.IsSameTree(tree1, tree2)
			if result != tt.expected {
				t.Errorf("IsSameTree() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
