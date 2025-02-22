package day8

// Time Complexity: O(N)
// Space Complexity: O(1)

// 雙指針
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
