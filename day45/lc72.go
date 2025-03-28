package day45

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以i-1為結尾的word1與j-1為結尾的word2為相同的最小操作次數
// 遞推公式:
// 相同: dp[i-1][j-1]
// 不相同:
// (刪除/增加)(兩者操作是相同的) dp[i-1][j]+1, dp[i][j-1]+1
// 替換 dp[i-1][j-1]+1
// min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)

// dp初始化:
// dp[i][0] = i, dp[0][j] = j
// 遍歷順序: 前到後, 上到下

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m + 1 {
		dp[i][0] = i
	}

	for j := range n + 1 {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}
