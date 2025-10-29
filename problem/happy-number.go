package problem

// https://leetcode.com/problems/happy-number/

// IsHappy determines if a number is a happy number.
// A happy number is defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits,
func IsHappy(n int)bool {
	seen := make(map[int]bool)

	for {
		if n == 1 {
			return true
		}
		if seen[n] {
			return false
		}
		seen[n] = true
		n = sumOfSquares(n)
	}	
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
