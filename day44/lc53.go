package day44

// Time Complexity: O(n)
// Space Complexity: O(n)

// dp[i][j]: 以i結尾的最大連續子序列和
// 遞推公式: max(dp[i-1]+nums[i], nums[i])
// dp初始化: nums[0]
// 遍歷順序: 前往後

// 有貪心解
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := nums[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}
