package problem

// https://leetcode.com/problems/next-greater-numerically-balanced-number

func NextGreaterNumber(n int) int {
	for x := n + 1; ; x++ {
		if isNumericallyBalanced(x) {
			return x
		}	
	}
}

func isNumericallyBalanced(num int) bool {
	var count [10]int
	if num == 0 {
		count[0] = 1
	}
	for v := num; v > 0; v /= 10 {
		count[v%10]++
	}

	// digit 0 must not appear
	if count[0] > 0 {
		return false
	}
	// for digits 1-9, either does not appear
	// or appears exactly d times
	for d := 1; d <= 9; d++ {
		if count[d] != 0 && count[d] != d {
			return false
		}
	}
	return true
}
