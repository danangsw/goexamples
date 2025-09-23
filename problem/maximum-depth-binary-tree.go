// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package problem

import helper "../helper"

// import "runtime/debug"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// MaxDepthRecursive finds the maximum depth of a binary tree using recursive DFS.
// This is the most elegant and intuitive solution.
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(h) where h is the height of the tree (recursion stack)
// Best For: Most interviews and general use cases
func MaxDepthRecursive(root *TreeNode) int {
	// Base case: empty tree has depth 0
	if root == nil {
		return 0
	}

	// Recursive case: 1 + maximum depth of left and right subtrees
	leftDepth := MaxDepthRecursive(root.Left)
	rightDepth := MaxDepthRecursive(root.Right)

	// Return 1 (current node) + maximum of left and right depths
	if leftDepth > rightDepth {
		return 1 + leftDepth
	}
	return 1 + rightDepth
}

// MaxDepthRecursiveOneLiner is the most concise version using built-in max function.
// This demonstrates Go's functional programming capabilities.
func MaxDepthRecursiveOneLiner(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + helper.Max(MaxDepthRecursiveOneLiner(root.Left), MaxDepthRecursiveOneLiner(root.Right))
}

// MaxDepthIterativeBFS finds the maximum depth using iterative BFS (level-order traversal).
// This approach processes the tree level by level.
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(w) where w is the maximum width of the tree
// Best For: When you need level-by-level processing or want to avoid recursion
func MaxDepthIterativeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		// Process all nodes at current level
		levelSize := len(queue)
		depth++

		// Process each node at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // Get first element
			queue = queue[1:] // Remove first element (dequeue)

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// MaxDepthIterativeDFS finds the maximum depth using iterative DFS with a stack.
// This approach simulates the recursive call stack manually.
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(h) where h is the height of the tree
// Best For: When recursion is not preferred but DFS traversal is needed
func MaxDepthIterativeDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Stack to store pairs of (node, depth)
	type nodeDepth struct {
		node  *TreeNode
		depth int
	}

	stack := []nodeDepth{{root, 1}}
	maxDepth := 0

	for len(stack) > 0 {
		// Pop from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node, depth := current.node, current.depth

		// Update maximum depth
		if depth > maxDepth {
			maxDepth = depth
		}

		// Push children with incremented depth
		if node.Left != nil {
			stack = append(stack, nodeDepth{node.Left, depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, nodeDepth{node.Right, depth + 1})
		}
	}

	return maxDepth
}

// MaxDepthIterativeDFSOptimized is an optimized version that tracks depth without storing it.
// This version uses less memory by tracking depth separately.
func MaxDepthIterativeDFSOptimized(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeStack := []*TreeNode{root}
	depthStack := []int{1}
	maxDepth := 0

	for len(nodeStack) > 0 {
		// Pop from both stacks
		node := nodeStack[len(nodeStack)-1]
		depth := depthStack[len(depthStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		depthStack = depthStack[:len(depthStack)-1]

		// Update maximum depth
		if depth > maxDepth {
			maxDepth = depth
		}

		// Push children with incremented depth
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			depthStack = append(depthStack, depth+1)
		}
		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			depthStack = append(depthStack, depth+1)
		}
	}

	return maxDepth
}

func init() {
	//debug.SetMemoryLimit(2 * 1024 * 1024) // Limit memory usage to 2MB for testing purposes
}
