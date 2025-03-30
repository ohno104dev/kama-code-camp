package day46

// Time Complexity: O(n^2)
// Space Complexity: O(n^2)

// dp[i][j]: [i, j]的迴文子序列長度
// 遞推公式:
// 相同 dp[i+1][j-1]+2
// 不相同 max(dp[i][j-1], dp[i+1][j])
// dp初始化: dp[i:i] = 1 , 單一字符
// 遍歷順序: 前到後, 下到上

func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for i := range size {
		dp[i] = make([]int, size)
		dp[i][i] = 1 // 單一字符
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][size-1]
}
