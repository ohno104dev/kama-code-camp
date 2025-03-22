package day35

// goos:
// weight:[ 1,  3,  4]
// value: [15, 20, 30]
// find max value at bag weight = 4

// Time Complexity: O(n*w), n個物品 * w背包容量
// Space Complexity: O(n*w)

// dp[i][j]: [0,i]的物品任取放進容量為j的背包
// 遞推公式: dp[i][j] = max( dp[i-1][j](不放物品i), dp[i-1][j-weight[i]]+value[i](放物品i) )
// dp初始化: dp[0][j] = 物品0的value, dp[i][0] = 0
// 遍歷順序: 左往右, 上往下

// 2維數組解
func bagProblem2D() int {
	goosWeight := []int{1, 3, 4}
	goosValue := []int{15, 20, 30}
	bagWeight := 4

	// 				背包重量j
	// 			0	1	2	3	4
	// 物品	0
	// 物品	1
	// 物品 2

	dp := make([][]int, len(goosWeight))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1) // 多加0, 從0開始為5
	}

	// 初始化第一件物品
	for j := 0; j <= bagWeight; j++ {
		if j >= goosWeight[0] {
			dp[0][j] = goosValue[0]
		}
	}

	// 先遍歷哪個不影響
	for i := 1; i < len(goosWeight); i++ { // 遍歷物品
		for j := 0; j <= bagWeight; j++ { // 遍歷背包容量
			if j < goosWeight[i] {
				dp[i][j] = dp[i-1][j] // 不放物品i
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-goosWeight[i]]+goosValue[i])
			}
		}
	}

	return dp[len(goosWeight)-1][bagWeight]
}
