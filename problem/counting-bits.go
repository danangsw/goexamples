package problem

// https://leetcode.com/problems/counting-bits/

// Shift-based approach
func CountBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// i >> 1 removes the least significant bit. 
		// If i is even last bit is 0; if odd last bit is 1.
		// So we add (i & 1) to account for that.

		// Example trace (shift-based for n=5):
		// i=1: ans[1]=ans[0]+1=1
		// i=2: ans[2]=ans[1]+0=1
		// i=3: ans[3]=ans[1]+1=2
		// i=4: ans[4]=ans[2]+0=1
		// i=5: ans[5]=ans[2]+1=2
		// Result: [0,1,1,2,1,2]
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

// Dynamic Programming approach
func CountBitsDP(n int) []int {
	result := make([]int, n+1)
	// offset tracks the largest power of 2 less than or equal to i
	offset := 1
	for i := 1; i <= n; i++ {
		// When i reaches the next power of 2, update offset
		if i == offset*2 {
			offset *= 2
		}
		// The number of 1's in i is 1 (for the leading 1) plus
		// the number of 1's in the remainder (i - offset)
		// Example trace (DP-based for n=5):
		// i=1: ans[1]=ans[0]+1=1
		// i=2: ans[2]=ans[0]+1=1
		// i=3: ans[3]=ans[1]+1=2
		// i=4: ans[4]=ans[0]+1=1
		// i=5: ans[5]=ans[1]+1=2
		// Result: [0,1,1,2,1,2]
		result[i] = result[i-offset] + 1
	}
	return result
}

// Modulo-based approach
func CountBitsModulo(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// i / 2 shifts the bits to the right (removes least significant bit)
		// i % 2 gives 1 if the least significant bit is 1, else 0
		// Example trace (modulo-based for n=5):
		// i=1: ans[1]=ans[0]+1=1
		// i=2: ans[2]=ans[1]+0=1
		// i=3: ans[3]=ans[1]+1=2
		// i=4: ans[4]=ans[2]+0=1
		// i=5: ans[5]=ans[2]+1=2
		// Result: [0,1,1,2,1,2]
		result[i] = result[i/2] + (i % 2)
	}
	return result
}