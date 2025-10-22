package problem

// https://leetcode.com/problems/counting-bits/

func CountBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// i >> 1 removes the least significant bit. 
		// If i is even last bit is 0; if odd last bit is 1.
		// So we add (i & 1) to account for that.
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}