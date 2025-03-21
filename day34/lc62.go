package day34

// Time Complexity: O(m*n)
// Space Complexity: O(m*n)

// 確認dp數組含義: 從(0,0)開始走到(i,j)有幾種
// 遞推公式: dp[i][j] = dp[i-1][j]+dp[i][j-1]
// dp數組如何初始化: dp[0][j], dp[i][0]皆為1種(題意)
// 遍歷順序: 左往右, 上往下
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
