package day35

// goos:
// weight:[ 1,  3,  4]
// value: [15, 20, 30]
// find max value at bag weight = 4

// Time Complexity: O(n*w), n個物品 * w背包容量
// Space Complexity: O(n*w)

// dp[j]: 容量為j的背包所能裝的最大value
// 遞推公式: dp[j] = max( dp[j](不放物品i), dp[j-weight[i]]+value[i](放物品i) )
// dp初始化: dp[0] = 0
// 遍歷順序: 物品前到後, 容量大到小

// 1維數組解 - 滾動數組, 只保留上一層的計算結果
func bagProblem1D() int {
	goosWeight := []int{1, 3, 4}
	goosValue := []int{15, 20, 30}
	bagWeight := 4

	// 				背包容量j
	// 			0	1	2	3	4
	// 最大價值  0

	dp := make([]int, bagWeight+1) // 多加0, 從0開始為5

	// 初始化
	dp[0] = 0 // 容量為0的背包所能裝的最大value

	// 只能先遍歷物品
	for i := range goosWeight { // 遍歷物品
		// 倒敘遍歷背包容量, 避免重複選取
		for j := bagWeight; j >= goosWeight[i]; j-- {
			dp[j] = max(dp[j], dp[j-goosWeight[i]]+goosValue[i])
		}
	}

	return dp[bagWeight]
}
