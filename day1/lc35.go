package day1

// Time Complexity: O(log(N))
// Space Complexity: O(1)

// 使用二分查找(704)概念
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) // 左閉右開

	for left < right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid
		}
	}

	// left最後會等於mid
	// 最小mid index = 0
	// 最大mid index = len(nums)+1

	return left
}
