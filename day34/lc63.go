package day34

// Time Complexity: O(m*n)
// Space Complexity: O(m*n)

// dp[i][j]: 從(0,0)開始走到(i,j)有幾種
// 遞推公式: dp[i][j] = dp[i-1][j]+dp[i][j-1]
// dp初始化: dp[0][j], dp[i][0]遇到障礙前皆為1種(題意)
// 遍歷順序: 左往右, 上往下
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
