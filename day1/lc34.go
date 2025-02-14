package day1

// Time Complexity: O(log(N)+log(N))
// Space Complexity: O(1)

// 使用二分查找(704)概念
func searchRange(nums []int, target int) []int {
	// 會出現[3], target=3, ans = {0, 0}的情況
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

func searchFirstEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1

		if target == nums[mid] {
			// 會出現[3,3,3,3,3,3]的情況, 必須找到最前一個
			if mid == 0 || target != nums[mid-1] {
				return mid
			}
			right = mid - 1
		}

		if target > nums[mid] {
			left = mid + 1
		}

		if target < nums[mid] {
			right = mid - 1
		}
	}

	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1

		if target == nums[mid] {
			// 會出現[3,3,3,3,3,3]的情況, 必須找到最後一個
			if mid == len(nums)-1 || target != nums[mid+1] {
				return mid
			}
			left = mid + 1
		}

		if target > nums[mid] {
			left = mid + 1
		}

		if target < nums[mid] {
			right = mid - 1
		}
	}

	return -1
}
