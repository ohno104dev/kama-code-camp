package day36

// Time Complexity: O(m*n)
// Space Complexity: O(n)

// dp[i][j]: i個0, j個1, 最大裝[i][j]個物品
// 遞推公式: max(dp[i-x][j-y]+1, dp[i][j])
// dp初始化: dp[0][0] = 0
// 遍歷順序: 物品前到後, 容量大到小

// 二維0/1背包
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range strs {
		zeroCount, oneCount := 0, 0
		for _, v := range strs[i] {
			if v == '0' {
				zeroCount++
			}
			if v == '1' {
				oneCount++
			}
		}

		for x := m; x >= zeroCount; x-- {
			for y := n; y >= oneCount; y-- {
				dp[x][y] = max(dp[x-zeroCount][y-oneCount]+1, dp[x][y])
			}
		}
	}

	return dp[m][n]
}
