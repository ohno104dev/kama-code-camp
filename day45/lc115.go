package day45

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以i-1為結尾的s中有以j-1為結尾t的s個數
// 遞推公式:
// 相同 dp[i-1][j-1]+ dp[i-1][j]
// 不相同 dp[i-1][j]
// dp初始化:
// dp[i][0] = 1, dp[0][j] = 0, dp[0][0] = 1
// 遍歷順序: 前到後, 上到下

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := range dp {
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
