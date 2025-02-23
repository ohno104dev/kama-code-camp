package day9

// Time Complexity: O(N+M)
// Space Complexity: O(N)

// KMP
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j前一位的匹配點
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

// 看前一位
func getNext(next []int, s string) {
	j := 0
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

// // Sample Solution
// // Time Complexity: O(N*M)
// func strStr(haystack string, needle string) int {
// 	if needle == "" {
// 		return 0
// 	}
// 	h := len(haystack)
// 	n := len(needle)
// 	for i := 0; i < h-n+1; i++ {
// 		if haystack[i:i+n] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }
