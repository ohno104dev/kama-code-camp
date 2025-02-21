package day6

// Time Complexity: O(N)
// Space Complexity: O(N)

// 注意 string, char, rune
func isAnagram(s string, t string) bool {
	record := [26]int{} // 固定範圍

	for _, r := range s {
		record[r-'a']++
	}

	for _, r := range t {
		record[r-'a']--
	}

	return record == [26]int{} // record如果有不为零0,則不相等
}
