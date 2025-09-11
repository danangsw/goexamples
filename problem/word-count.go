package problem

//import "fmt"

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	if len(s) == 0 {
		return result
	}

	return result
}

// s: Hello World Let's Go!
// del: ' '
// [Hello World Let's Go!]
func StringSplit(s string, del rune) []string {
	var result = []string{}
	if len(s) == 0 {
		return result
	}
	start := 0
	for i, r := range s {
		//fmt.Println(i, start, r, del, string(r), string(del))
		if r == del && i >= start {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	if start <= len(s) {
		result = append(result, s[start:])
	}
	return result
}
