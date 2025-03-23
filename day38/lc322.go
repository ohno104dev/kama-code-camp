package day38

import "math"

// Time Complexity: O(n*amount) // n是coins的長度
// Space Complexity: O(amount)

// dp[j]: 裝滿容量j的背包, 最少為dp[j]
// 遞推公式: min(dp[j], dp[j-coins[i]]+1)
// dp初始化: dp[0]=0 , 非零值為最大值 (求最小)
// 遍歷順序: 求各數, 皆可

// 完全背包, 前到後
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 先遍歷物品再遍歷背包, 組合數
	// 先遍歷背包再遍歷物品, 排列數
	// 求各數, 皆可
	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
