package day9

// Time Complexity: O(N)
// Space Complexity: O(N)

// 利用KMP(28)的前綴表概念
// 數學證明
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if len(s) == 0 {
		return false
	}

	next := make([]int, n)
	j := 0
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}

	// 原字符串長度 % (原字符串長度-最長相等前後綴的長度) == 0, 表示由重複子字串組成
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}

	return false
}

// // 移動匹配解
// // 數學證明
// func repeatedSubstringPattern(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}
// 	t := s + s

// 	return strings.Contains(t[1:len(t)-1], s)
// }
