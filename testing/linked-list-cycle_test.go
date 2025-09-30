package testing_test

import (
    "testing"
	"fmt"
    helper "../helper"
    problem "../problem"
)

func TestHasCycle(t *testing.T) {
    cases := []struct {
        name     string
        head     *helper.ListNode
        expected bool
    }{
        {
            name:     "Empty list",
            head:     nil,
            expected: false,
        },
        {
            name:     "Single node no cycle",
            head:     helper.BuildLinkedList([]int{1}),
            expected: false,
        },
        {
            name:     "Single node cycle",
            head:     createCycleList([]int{1}, 0), // Cycle to itself
            expected: true,
        },
        {
            name:     "Two nodes no cycle",
            head:     helper.BuildLinkedList([]int{1, 2}),
            expected: false,
        },
        {
            name:     "Two nodes with cycle",
            head:     createCycleList([]int{1, 2}, 0), // 2 -> 1
            expected: true,
        },
        {
            name:     "No cycle - linear list",
            head:     helper.BuildLinkedList([]int{1, 2, 3, 4, 5}),
            expected: false,
        },
        {
            name:     "Cycle at head",
            head:     createCycleList([]int{1, 2, 3, 4, 5}, 0), // 5 -> 1
            expected: true,
        },
        {
            name:     "Cycle in middle",
            head:     createCycleList([]int{1, 2, 3, 4, 5}, 2), // 5 -> 3
            expected: true,
        },
        {
            name:     "Cycle at end",
            head:     createCycleList([]int{1, 2, 3, 4, 5}, 4), // 5 -> 5
            expected: true,
        },
        {
            name:     "Large list no cycle",
            head:     helper.BuildLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
            expected: false,
        },
        {
            name:     "Large list with cycle",
            head:     createCycleList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), // 10 -> 6
            expected: true,
        },
        {
            name:     "Cycle to second node",
            head:     createCycleList([]int{1, 2, 3, 4, 5}, 1), // 5 -> 2
            expected: true,
        },
        {
            name:     "Cycle to fourth node",
            head:     createCycleList([]int{1, 2, 3, 4, 5}, 3), // 5 -> 4
            expected: true,
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            result := problem.HasCycle(tc.head)
            if result != tc.expected {
                t.Errorf("HasCycle() = %v; expected %v", result, tc.expected)
            }
        })
    }
}

// Helper function to create a linked list with a cycle
// cycleIndex: index where the last node should point to (0-based)
// If cycleIndex == -1, no cycle is created
func createCycleList(values []int, cycleIndex int) *ListNode {
    if len(values) == 0 {
        return nil
    }

    // Create the linear list
    head := helper.BuildLinkedList(values)

    if cycleIndex == -1 {
        return head // No cycle
    }

    // Find the node at cycleIndex
    cycleNode := head
    for i := 0; i < cycleIndex && cycleNode != nil; i++ {
        cycleNode = cycleNode.Next
    }

    // Find the last node
    lastNode := head
    for lastNode.Next != nil {
        lastNode = lastNode.Next
    }

    // Create the cycle
    if cycleNode != nil {
        lastNode.Next = cycleNode
    }

    return head
}

func TestHasCycleEdgeCases(t *testing.T) {
    t.Run("Nil head", func(t *testing.T) {
        result := problem.HasCycle(nil)
        if result != false {
            t.Errorf("HasCycle(nil) = %v; expected false", result)
        }
    })

    t.Run("Single node pointing to itself", func(t *testing.T) {
        node := &ListNode{Val: 1}
        node.Next = node // Cycle to itself
        result := problem.HasCycle(node)
        if result != true {
            t.Errorf("HasCycle(single cycle) = %v; expected true", result)
        }
    })

    t.Run("Two nodes cycle back to first", func(t *testing.T) {
        node1 := &ListNode{Val: 1}
        node2 := &ListNode{Val: 2}
        node1.Next = node2
        node2.Next = node1 // Cycle: 1 -> 2 -> 1
        result := problem.HasCycle(node1)
        if result != true {
            t.Errorf("HasCycle(two node cycle) = %v; expected true", result)
        }
    })
}

func BenchmarkHasCycle(b *testing.B) {
    // Benchmark with different list sizes
    sizes := []int{100, 1000, 10000}

    for _, size := range sizes {
        b.Run(fmt.Sprintf("no_cycle_size_%d", size), func(b *testing.B) {
            // Create linear list
            head := helper.BuildLinkedList(make([]int, size))
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                problem.HasCycle(head)
            }
        })

        b.Run(fmt.Sprintf("with_cycle_size_%d", size), func(b *testing.B) {
            // Create list with cycle
            head := createCycleList(make([]int, size), size/2)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                problem.HasCycle(head)
            }
        })
    }
}
