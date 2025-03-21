package day34

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// 確認dp數組含義: 對i進行拆分,最大乘積
// 遞推公式: max((i-j)*j, dp[i-j]*j, dp[i])
// dp數組如何初始化: dp[0]=0, dp[1]=0, dp[2]=1
// 遍歷順序: 從小到大
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], (i-j)*j, dp[i-j]*j)
		}
	}

	return dp[n]
}
