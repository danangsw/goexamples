// https://leetcode.com/problems/roman-to-integer/description/

package problem

func RomanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	lens := len(s)
	total := romans[s[lens-1]]

	for i := lens-2; i >= 0; i-- {
		if romans[s[i]] < romans[s[i+1]] {
			total -= romans[s[i]]
		} else {
			total += romans[s[i]]
		}
	}

	return total
}
