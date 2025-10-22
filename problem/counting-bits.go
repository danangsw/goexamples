package problem

// https://leetcode.com/problems/counting-bits/

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