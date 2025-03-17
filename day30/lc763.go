package day30

// Time Complexity: O(N)
// Space Complexity: O(1)

func partitionLabels(s string) []int {
	marks := make(map[byte]int, 26)

	// 找出每個字符最遠的位置
	for i := range len(s) {
		marks[s[i]] = i
	}

	res := []int{}
	left, right := 0, 0
	for i := range len(s) {
		right = max(right, marks[s[i]])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return res
}
