package problem

// https://leetcode.com/problems/majority-element/description/
// Boyer-Moore Voting Algorithm (Optimal Solution)

func MajorityElement(nums []int) int {
	counter := 1
	candidate := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			candidate = nums[i]
			counter = 1
		}
	}
	return candidate
}
