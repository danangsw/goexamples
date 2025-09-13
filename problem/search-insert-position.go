package problem

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for i := 0; i < len(nums)-1; i++ {
		if target <= nums[i] {
			return i
		}
		if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}

	return len(nums)
}
