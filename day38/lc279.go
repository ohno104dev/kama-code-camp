package day38

import "math"

// Time Complexity: O(n*√n)
// Space Complexity: O(n)

// dp[j]: 和為j的完全平方數的最少數量為dp[j]
// 遞推公式: min(dp[j], dp[j-i*i]+1)
// dp初始化: dp[0]=0 , 非零值為最大值 (求最小)
// 遍歷順序: 求各數, 皆可

// 完全背包, 前到後, 使用(322)概念
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	// 先遍歷物品再遍歷背包, 組合數
	// 先遍歷背包再遍歷物品, 排列數
	// 求各數, 皆可
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}
