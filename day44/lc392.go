package day44

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以i-1為結尾的字串s與j-1為解為的字串t的公共相等子序列的最大長度
// 遞推公式:
// 相同 dp[i-1][j-1]+1
// 不相同 dp[i][j-1]
// dp初始化: dp[:] = 0
// 遍歷順序: 前到後, 上到下

// 使用(1143)概念
// 可以用雙指針
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}
