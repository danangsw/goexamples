// https://leetcode.com/problems/climbing-stairs/
package problem

// ClimbStairs calculates the number of distinct ways to climb
// to the top of a staircase with n steps, where each time you
// can either climb 1 or 2 steps. This is a classic dynamic programming
// problem that can be solved using an iterative approach with
// O(n) time complexity and O(1) space complexity.
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0 // or handle as appropriate
	}
	if n <= 2 {
		return n // n == 1 -> 1 way, n == 2 -> 2 ways
	}

	pos1 := 2      // ways to reach (i-1)th step
	pos2 := 1      // ways to reach (i-2)th step
	var result int // ways to reach current step
	// Iteratively compute the number of ways to reach each step
	// using the relation: ways(i) = ways(i-1) + ways(i-2)
	// This is similar to Fibonacci sequence
	for i := 3; i <= n; i++ {
		result = pos1 + pos2 // Current ways is sum of the two previous
		pos2 = pos1          // Move pos2 to pos1 for next iteration
		pos1 = result        // Move pos1 to current result for next iteration
	}
	return result
}
