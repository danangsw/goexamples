// https://leetcode.com/problems/excel-sheet-column-number/description/
package problem

func TitleToNumber(columnTitle string) int {
	total := 0

	for _, char := range columnTitle {
		digit := int(char - 'A' + 1)
		total = total*26+digit
	}

	return total
}

// recursive solution
func TitleToNumberRecursive(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}
	lastChar := columnTitle[len(columnTitle)-1]
	digit := int(lastChar - 'A' + 1)
	return digit + 26*TitleToNumberRecursive(columnTitle[:len(columnTitle)-1])
}