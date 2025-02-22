package day7

// Time Complexity: O(N^2+N^2)
// Space Complexity: O(N)

func canConstruct(ransomNote string, magazine string) bool {
	// rec := make([]int, 26)
	rec := [26]int{} // 固定資料集

	for _, v := range magazine {
		rec[v-'a']++
	}

	for _, v := range ransomNote {
		rec[v-'a']--
		if rec[v-'a'] < 0 {
			return false
		}
	}
	return true
}
