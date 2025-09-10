// https://leetcode.com/problems/remove-duplicates-from-sorted-array

package problem

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var p int
	for i := range nums {
		if nums[p] != nums[i] {
			p++
			nums[p] = nums[i]
		}
	}
	p++
	return p
}
