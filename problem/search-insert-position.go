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

func SearchInsertBinary(nums []int, target int) int {
	// Initialize low and high: Set low to 0 and high to len(nums) - 1.
	low := 0
	high := len(nums) - 1

	// Iterate while low <= high:
	for low <= high {
		// Calculate the middle index (This prevents potential overflow).
		mid := low + (high-low)/2
		if nums[mid] == target {
			// Target found
			return mid
		}
		if nums[mid] < target {
			//  The target must be in the right half of the array
			low = mid + 1
		}
		if nums[mid] > target {
			//  The target must be in the left half of the array
			high = mid - 1
		}
	}

	return low
}
