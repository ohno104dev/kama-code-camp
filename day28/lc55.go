package day28

// Time Complexity: O(N)
// Space Complexity: O(1)

// 轉換思路, 只管能不能cover到最後一個位置
func canJump(nums []int) bool {
	cover := 0

	if len(nums) == 1 {
		return true
	}

	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}
