package day1

// Time Complexity: O(N)
// Space Complexity: O(N)

// 雙指針
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums)) // 給長度才能下標索引
	k := len(nums) - 1

	for left, right := 0, len(nums)-1; left <= right; {
		if nums[right]*nums[right] > nums[left]*nums[left] {
			res[k] = nums[right] * nums[right]
			right--
		} else {
			res[k] = nums[left] * nums[left]
			left++
		}
		k--
	}
	return res
}
