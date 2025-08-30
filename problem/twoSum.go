//https://leetcode.com/problems/two-sum/

package problem

func TwoSum(nums []int, target int) []int {
	numIdx := make(map[int]int) // Key: Value Value: Index

	for i, num := range nums {
		checkN := target - num
		if ix, found := numIdx[checkN]; found {
			return []int{i, ix}
		}

		numIdx[num] = i
	}

	return nil
}
