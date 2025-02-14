package day1

// Time Complexity: O(N)
// Space Complexity: O(N)

// 快慢指針
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow // index位置等於結果集的長度
}
