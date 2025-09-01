package samples

import (
	"fmt"
	"strings"
)

type FilterCallback func(string) bool

func FunctionExample(contain string, prefix string, length int) {
	fmt.Println("Function Examples")

	// Example usage of filter function
	data := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
		"pinnaples",
		"watermelon",
		"blueberry",
		"blackberry",
		"strawberry",
	}

	length1, result1 := filter(data, func(item string) bool {
		return strings.Contains(item, contain)
	})
	length2, result2 := filter(data, func(item string) bool {
		return strings.HasPrefix(item, prefix)
	})
	length3, result3 := filter(data, func(item string) bool {
		return len(item) >= length
	})

	fmt.Printf("Contains '%s': %d %v. \n", contain, length1, result1)
	fmt.Printf("Starts with '%s': %d %v. \n", prefix, length2, result2)
	fmt.Printf("Longer than %d characters: %d %v. \n", length, length3, result3)
}

func filter(data []string, callback FilterCallback) (length int, result []string) {
	for _, item := range data {
		if callback(item) {
			result = append(result, item)
		}
	}
	length = len(result)
	return
}
