package day44

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以text1[0, i-1]和text2[0, j-1]的最長公共子序列長度 (可以避免初始化dp[0][j]與dp[i][0])
// 遞推公式:
// 相同 dp[i][j]=dp[i-1][j-1]+1
// 不相同 max(dp[i-1][j], dp[i][j-1])
// dp初始化: dp[:] = 0
// 遍歷順序: 前到後, 上到下

func longestCommonSubsequence(text1 string, text2 string) int {
	t1 := len(text1)
	t2 := len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}
