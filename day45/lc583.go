package day45

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以i-1為結尾的word1與j-1為結尾的word2為相同的最小操作次數
// 遞推公式:
// 相同: dp[i-1][j-1]
// 不相同: min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
// dp初始化:
// dp[i][0] = i, dp[0][j] = j
// 遍歷順序: 前到後, 上到下

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := range dp {
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// // 使用(1143)概念
// func minDistance(word1 string, word2 string) int {
// 	dp := make([][]int, len(word1)+1)
// 	for i := range dp {
// 		dp[i] = make([]int, len(word2)+1)
// 	}
// 	for i := 1; i <= len(word1); i++ {
// 		for j := 1; j <= len(word2); j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}

// 	return len(word1) + len(word2) - dp[len(word1)][len(word2)]*2
// }
