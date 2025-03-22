package day35

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// dp[j]: 容量為j的最大值
// 遞推公式: max(dp[j], dp[j-nums[i]]+nums[i]) (題意, 重量與價值相同)
// dp初始化: dp[0]=0
// 遍歷順序: 物品前到後, 容量大到小

// 轉換思路0/1背包
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 	不能對分
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	for i := range nums {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}
