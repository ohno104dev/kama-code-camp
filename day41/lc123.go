package day41

// Time Complexity: O(n)
// Space Complexity: O(n*5)

// []dp{}: 返回數組dp[0]:持有股票最大現金值, dp[1]:不持有股票最大現金值
// 遞推公式:
// 沒有操作		   dp[0][0] = 0
// 第一次持有股票	dp[0][1] = -prices[0]
// 第一次不持有股票	dp[0][2] = 0
// 第二次持有股票	dp[0][3] = -prices[0]
// 第二次不持有股票	dp[0][4] = 0
// dp初始化: dp[0][0] = 0
// 遍歷順序: 小到大

// 二次買賣
func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range prices {
		dp[i] = make([]int, 5)
	}

	dp[0][0] = 0          // 沒有操作
	dp[0][1] = -prices[0] // 第一次持有股票
	dp[0][2] = 0          // 第一次不持有股票
	dp[0][3] = -prices[0] // 第二次持有股票
	dp[0][4] = 0          // 第二次不持有股票

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]
}
