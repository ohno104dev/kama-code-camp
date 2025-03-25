package day42

// Time Complexity: O(n)
// Space Complexity: O(n)

// []dp{}:
// dp[i][0]: 持有股票
// dp[i][1]: 保持賣出股票
// dp[i][2]: 賣出股票
// dp[i][3]: 冷凍期

// 遞推公式:
// dp[i][0]:
// 前一天持有: dp[i-1][0]
// 前一天冷凍期買入股票: dp[i-1][3]-prices[i]
// 前一天保持賣出: dp[i-1][1]-prices[i]

// dp[i][1]:
// 前一天保持賣出: dp[i-1][1]
// 前一天冷凍期: dp[i-1][3]

// dp[i][2]:
// 前一天持有: dp[i-1][0]+prices[i]

// dp[i][3]:
// 前一天賣出: dp[i-1][2]

// dp初始化:
// dp[0][0]: -prices[0]
// dp[0][1]: 0 (非法狀態, 根據題意判斷)
// dp[0][2]: 0 (非法狀態, 根據題意判斷)
// dp[0][3]: 0

// 遍歷順序: 小到大

// 使用(122)概念
func maxProfit5(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], dp[n-1][2], dp[n-1][3])
}
