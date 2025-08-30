//https://leetcode.com/problems/palindrome-number/

package problem

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	reversed := 0
	original := x

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == original
}
