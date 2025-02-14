package day1

// Time Complexity: O(log(N))
// Space Complexity: O(1)

// 左閉右開
func search(nums []int, target int) int {
	left := 0
	right := len(nums) // 左閉右開

	for left < right { // 左閉右開
		mid := left + (right-left)>>1 // (防止溢出) 右移,除2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid // 左閉右開
		}

	}

	return -1
}

// 左閉右閉
// func search(nums []int, target int) int {
// 	left := 0
// 	right := len(nums) - 1 // 左閉右閉

// 	for left <= right { //左閉右閉
// 		middle := left + (right-left)>>1 // (prevent overflow) 右移,除2

// 		if target == nums[middle] {
// 			return middle
// 		} else if target > nums[middle] {
// 			left = middle + 1
// 		} else {
// 			right = middle - 1 // 左閉右閉
// 		}
// 	}

// 	return -1
// }
