package day31

import "strconv"

// Time Complexity: O(N)
// Space Complexity: O(N)

// 注意轉換處理
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	flag := -1

	// 處理數字的方向跟slice方向相反
	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			flag = i
			s[i-1]--
		}
	}

	if flag >= 0 {
		for i := flag; i < len(s); i++ {
			s[i] = '9'
		}
	}

	num, _ := strconv.Atoi(string(s))
	return num
}
