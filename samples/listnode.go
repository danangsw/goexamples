package samples

import "fmt"

type ListNode struct {
	Value int       // The value stored in the node
	Next  *ListNode // Pointer to the next node
}

func ListNodeExample() {
	basicListSample()
	node1 := createList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	printList(node1)
	fmt.Println("Create List")
	node2 := createList([]int{})
	printList(node2)
	fmt.Println("Insert List, before:")
	printList(node1)
	node3 := appendLists(node1, []int{11, 12, 13, 14, 15})
	fmt.Println("Insert List, after:")
	printList(node1)
	printList(node3)
}

func appendLists(nodes *ListNode, values []int) *ListNode {
	if len(values) == 0 || nodes == nil {
		return nil
	}

	// Find the latest node
	last := nodes
	for last.Next != nil {
		last = last.Next
	}

	// Append the values to List Node
	last.Next = createList(values)

	//fmt.Println(nodes, last)
	//printList(nodes)
	//printList(last)

	return nodes
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Value: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Value: values[i]}
		current = current.Next
	}

	return head
}

func basicListSample() {
	node1 := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	node3 := &ListNode{Value: 3}
	node4 := &ListNode{Value: 4}

	node1.Next = node2
	node2.Next = node3

	fmt.Println("Initial list")
	printList(node1)

	fmt.Println("Insert list")
	node2.Next = node4
	node4.Next = node3
	printList(node1)

	fmt.Println("Delete list")
	node2.Next = node3
	printList(node1)
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println(current)
}
