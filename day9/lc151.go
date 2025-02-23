package day9

// Time Complexity: O(N)
// Space Complexity: O(N)

// 雙指針
// 多反轉解題思路
func reverseWords(s string) string {
	b := []rune(s)

	slow := 0
	for fast := 0; fast < len(b); fast++ {
		// 處理思路不同, 判斷是字符開頭加入空白, 不是判斷空白本身
		if b[fast] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}

			for fast < len(b) && b[fast] != ' ' {
				b[slow] = b[fast]
				slow++
				fast++
			}
		}
	}

	b = b[:slow] // resize

	reverse(b)

	// reverse每個單詞
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}

	return string(b)
}

func reverse(s []rune) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
