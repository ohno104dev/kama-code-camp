package day32

// Time Complexity: O(N)
// Space Complexity: O(N)

// 確認dp數組含義: 達到i階有dp[i]種方法
// 遞推公式: dp[i-2]+dp[i-1]
// dp數組如何初始化: dp[0]無意義, dp[1]=1, dp[2]=2
// 遍歷順序: 前向後
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1, n+1) // 不使用dp[0]
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
