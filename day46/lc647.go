package day46

// Time Complexity: O(n^2)
// Space Complexity: O(n^2)

// dp[i][j]: [i,j]子串是否是回文子串
// 遞推公式:
// 單一字符 =>  true
// i j相差一 => true
// i j相差大於一, 判斷dp[i+1][j-1]是否迴文
// dp初始化: dp[:] = false
// 遍歷順序: 前到後, 下到上

// 可以用雙指針
func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := range len(s) {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}

	return res
}

// 雙指針
// func countSubstrings(s string) int {
// 	extend := func(i, j int) int {
// 		res := 0
// 		for i >= 0 && j < len(s) && s[i] == s[j] {
// 			i --
// 			j ++
// 			res ++
// 		}
// 		return res
// 	}
// 	res := 0
// 	for i := 0; i < len(s); i++ {
// 		res += extend(i, i)  // 以i为中心
// 		res += extend(i, i+1) // 以i和i+1为中心
// 	}
// 	return res
// }
