package day36

// Time Complexity: O(m*n)
// Space Complexity: O(n)

// dp[j]: 在容量為j時有幾種方法
// 遞推公式: dp[j]+dp[j-nums[i]]
// dp初始化: dp[0] = 0
// 遍歷順序: 物品前到後, 容量大到小

// 轉換思路0/1背包, 使用(1049)概念
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// P: 所有選擇+的數之和
	// N: 所有選擇-的數之和
	// sum = P + N
	// P - N = target

	// 目标值超出范围，无法找到解
	if target > sum || (sum+target)%2 == 1 || (sum+target)/2 < 0 {
		return 0
	}

	// P = (sum + target) / 2
	bag := (sum + target) / 2

	// DP 数组初始化
	dp := make([]int, bag+1)
	dp[0] = 1

	// 0-1 背包填充
	for _, num := range nums {
		for j := bag; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[bag]
}
