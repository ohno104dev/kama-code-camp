package day23

import "slices"

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n^2)

// 回溯
func partition(s string) [][]string {
	path := make([]string, 0)
	res := make([][]string, 0)

	var dfs func(s string, start int)

	dfs = func(s string, start int) {
		if start == len(s) {
			res = append(res, slices.Clone(path))
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				dfs(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(s, 0)

	return res
}

// 可以使用動態規劃優化回文串判斷
func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
