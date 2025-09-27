package problem

// https://leetcode.com/problems/pascals-triangle-ii/
func GetRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		// Update the row in reverse order, without using extra space
		for j := i; j > 0; j-- {
			// Update from the end to the beginning to avoid overwriting
			result[j] += result[j-1]
		}
	}
	return result
}
