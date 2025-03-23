package day36

// Time Complexity: O(m*n)
// Space Complexity: O(n)

// dp[j]: 在容量j的最大價值
// 遞推公式: max(dp[j], dp[j-stones[i]]+stones[i])
// dp初始化: dp[0] = 0
// 遍歷順序: 物品前到後, 容量大到小

// 轉換思路0/1背包, 使用(416)概念
func lastStoneWeightII(stones []int) int {

	dp := make([]int, 1501) //30*100/2+1
	dp[0] = 0

	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	for i := range stones {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - 2*dp[target]
}
