package day42

// Time Complexity: O(n)
// Space Complexity: O(n)

// []dp{}: 返回數組dp[0]:持有股票最大現金值, dp[1]:不持有股票最大現金值
// 遞推公式:
// 持有: max(dp[i-1][0], dp[i-1][1]-prices[i])
// 不持有: max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
// dp初始化: []dp{0,0}

// 遍歷順序: 小到大

// 使用(122)概念
func maxProfit2(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range prices {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] // 持有股票
	dp[0][1] = 0          // 沒持有股票

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}
