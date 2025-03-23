package day37

// Time Complexity: O(m*n) // m為amount, n是coins的長度
// Space Complexity: O(m)

// dp[j]: 裝滿容量j有幾種方法
// 遞推公式: dp[j]+=dp[j-coins[i]]
// dp初始化: dp[0]=1
// 遍歷順序:
// 求組合數, 先遍歷物品再遍歷背包

// 完全背包, 前到後
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	// 先遍歷物品再遍歷背包, 組合數
	// 先遍歷背包再遍歷物品, 排列數
	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
