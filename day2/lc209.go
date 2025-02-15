package day2

// Time Complexity: O(2N)
// Space Complexity: O(1)

// 滑動窗口
func minSubArrayLen(target int, nums []int) int {
	ans := len(nums) + 1 // 結果不會大於len(nums)
	start, sum := 0, 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
	}

	if ans > len(nums) {
		return 0
	}

	return ans
}
