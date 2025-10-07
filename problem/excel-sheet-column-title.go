package problem

// https://leetcode.com/problems/excel-sheet-column-title/description/

func ConvertToTitle(columnNumber int) string {
	result := ""
	return convertToTitleRecursive(columnNumber, result)
}

// recursive solution
func convertToTitleRecursive(columnNumber int, result string) string {
	if columnNumber == 0 {
		return result
	}
	columnNumber--
	rem := columnNumber % 26
	result = string(rune('A'+rem)) + result
	return convertToTitleRecursive(columnNumber/26, result)
}
