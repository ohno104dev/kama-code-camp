package day24

import (
	"strconv"
	"strings"
)

// Time Complexity: O(3^4)
// Space Complexity: O(n)

var (
	path []string
	res  []string
)

// 回溯
func restoreIpAddresses(s string) []string {
	path, res = make([]string, 0), make([]string, 0)
	ipdfs(s, 0)
	return res
}

func ipdfs(s string, start int) {
	if len(path) == 4 {
		if start == len(s) {
			ip := strings.Join(path, ".")
			res = append(res, ip)
		}
		return
	}

	// 依據每一組數值判斷
	for i := start; i < len(s); i++ {
		// 這組數字出現0開頭的無效數值
		if i != start && s[start] == '0' {
			break
		}

		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str)
			ipdfs(s, i+1)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}
