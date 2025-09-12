// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string

package problem

func FindString(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}
