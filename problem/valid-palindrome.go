package problem

// https://leetcode.com/problems/valid-palindrome/

import (
	"unicode"
)

func ValidPalindrome(s string) bool {
	// Convert string to rune slice to handle Unicode characters properly
	runes := []rune(s)
	left, right := 0, len(runes)-1

	// Use two-pointer technique to check for palindrome.
	// The left < right loop is tight and avoids unnecessary comparisons.
	for left < right {
		if !isAlphanum(runes[left]) {
			left++
			continue
		}
		if !isAlphanum(runes[right]) {
			right--
			continue
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphanum(c rune) bool {
	// Make simple check for alphanumeric characters
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
