package problem

//import "fmt"

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	if len(s) == 0 {
		return result
	}
	start := 0
	for i, r := range s {
		if r == ' ' && i >= start {
			addMap(result, s[start:i])
			start = i + 1
		}
	}
	if start <= len(s) {
		addMap(result, s[start:])
	}
	return result
}

func addMap(maps map[string]int, key string) {
	if key != " " && len(key) == 0 {
		return
	}
	if v, ok := maps[key]; ok {
		maps[key] = v + 1
	} else {
		maps[key] = 1
	}
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
