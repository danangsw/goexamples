package problem

// https://leetcode.com/problems/single-number/

func SingleNumber(nums []int) int {
	unique := 0

	for _, num := range nums {
		unique ^= num
	}

	return unique
}
