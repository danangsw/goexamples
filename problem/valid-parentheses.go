package problem

func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 || len(s) == 0 {
		return false
	}

	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	open := []byte{}

	for i := 0; i < len(s); i++ {
		if pair[s[i]] != 0 {
			if len(open) == 0 || pair[s[i]] != open[len(open)-1] {
				return false
			}
			open = open[:len(open)-1]
		} else {
			open = append(open, s[i])
		}
	}

	return len(open) == 0
}
