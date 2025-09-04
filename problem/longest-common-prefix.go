// https://leetcode.com/problems/longest-common-prefix/

package problem

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := range prefix {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != prefix[i] {
				return prefix[:i]
			}
		}
	}

	return prefix
}
