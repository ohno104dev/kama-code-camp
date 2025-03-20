package day32

// 確認dp數組含義: 到達i位置的最小花費為dp[i]
// 遞推公式: dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
// dp數組如何初始化: dp[0]=0, dp[1]=0 (題意)
// 遍歷順序: 前向後

// Time Complexity: O(N)
// Space Complexity: O(N)
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1, len(cost)+1)
	dp[0], dp[1] = 0, 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
