package testing_test

import (
	"testing"

	helper "../helper"
	problem "../problem"
)

// Helper function to create a linked list from slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// Helper function to convert linked list to slice for easy comparison
func listToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "Duplicates at beginning",
			input:    []int{1, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates at end",
			input:    []int{1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates in middle",
			input:    []int{1, 2, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Multiple duplicate groups",
			input:    []int{1, 1, 2, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complex pattern",
			input:    []int{1, 1, 1, 2, 3, 3, 4, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Two nodes with duplicates",
			input:    []int{1, 1},
			expected: []int{1},
		},
		{
			name:     "Large values",
			input:    []int{100, 100, 200, 300, 300, 400},
			expected: []int{100, 200, 300, 400},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input list
			inputList := createList(tt.input)

			// Call the function
			result := problem.DeleteDuplicates(inputList)

			// Convert result to slice
			resultSlice := listToSlice(result)

			// Compare with expected
			if !helper.SlicesEqual(resultSlice, tt.expected) {
				t.Errorf("DeleteDuplicates(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestDeleteDuplicatesEdgeCases(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		result := problem.DeleteDuplicates(nil)
		if result != nil {
			t.Errorf("DeleteDuplicates(nil) = %v, expected nil", result)
		}
	})
}

func TestDeleteDuplicatesStructureIntegrity(t *testing.T) {
	// Test that the function properly maintains list structure
	t.Run("List structure integrity", func(t *testing.T) {
		input := []int{1, 1, 2, 3, 3, 4}
		head := createList(input)

		result := problem.DeleteDuplicates(head)

		// Verify the result is properly linked
		expected := []int{1, 2, 3, 4}
		actual := listToSlice(result)

		if !helper.SlicesEqual(actual, expected) {
			t.Errorf("Structure integrity test failed: got %v, expected %v", actual, expected)
		}

		// Verify no cycles exist by checking if we can traverse to the end
		current := result
		count := 0
		maxNodes := len(expected) + 10 // Safety buffer

		for current != nil && count < maxNodes {
			current = current.Next
			count++
		}

		if count >= maxNodes {
			t.Error("Possible infinite loop detected in result list")
		}
	})
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small_10", 10},
		{"Medium_100", 100},
		{"Large_1000", 1000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// Create input with alternating duplicates
			input := make([]int, bm.size)
			for i := 0; i < bm.size; i++ {
				input[i] = i / 2 // Creates pairs of duplicates
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				list := createList(input)
				problem.DeleteDuplicates(list)
			}
		})
	}
}

/**
func ExampleDeleteDuplicates() {
	// Create a list: 1 -> 1 -> 2 -> 3 -> 3
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}

	// Remove duplicates
	result := problem.DeleteDuplicates(head)

	// Print result: should be 1 -> 2 -> 3
	current := result
	for current != nil {
		print(current.Val)
		if current.Next != nil {
			print(" -> ")
		}
		current = current.Next
	}
	println()
	// Output: 1 -> 2 -> 3
}
**/
