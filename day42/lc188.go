package day42

// Time Complexity: O(n*k) // n為prices長度
// Space Complexity: O(n*k)

// []dp{}: 返回數組dp[0]:持有股票最大現金值, dp[1]:不持有股票最大現金值
// 遞推公式:
// 沒有操作		   dp[0][0] = 0
// 第i天持有股票	dp[i][1] = dp[i-1][0]-prices[i]
// 第i天不持有股票	dp[i][1] = dp[i-1][1]
// dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])
// dp[i][2] = max(dp[i-1][1] + prices[i], dp[i-1][2])
// dp初始化: dp[0][0] = 0
// 遍歷順序: 小到大

// 使用(123)概念
// K次買賣
func maxProfit4(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}

	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}

	return dp[len(prices)-1][2*k]
}
