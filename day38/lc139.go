package day38

// Time Complexity: O(n^3) // substr返回子字串的副本是O(n)
// Space Complexity: O(n)

// dp[i]: 字符串的長度i, 能被dict組合
// 遞推公式: dp[j] && dict[s[j:i]]
// dp初始化: dp[0]=true
// 遍歷順序:
// 求排列數, 先遍歷背包再遍歷物品

// 完全背包, 前到後
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	// 先遍歷物品再遍歷背包, 組合數
	// 先遍歷背包再遍歷物品, 排列數
	for i := 1; i <= len(s); i++ {
		for j := range i {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
