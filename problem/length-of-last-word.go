// https://leetcode.com/problems/length-of-last-word
package problem

import "strings"

// Solution 1: O(N)
func LengthOfLastWord1(s string) int {
	return len(strings.Fields(s)[len(strings.Fields(s))-1])
}

// Solution 2:
func LengthOfLastWord2(s string) int {
	lens := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			lens++
		}
		if lens > 0 && s[i] == ' ' {
			return lens
		}
	}
	return lens
}
