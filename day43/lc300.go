package day43

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// []dp{}: 以nums[i]為結尾的最長遞增子序列的長度
// 遞推公式: max(dp[j]+1,dp[i])
// dp初始化: dp[:] = 1
// 遍歷順序: 前到後

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := range i {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}

	return ans
}
