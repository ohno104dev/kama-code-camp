package day29

// Time Complexity: O(N)
// Space Complexity: O(1)

// 轉換思路, 先確定一邊在確定另一邊, 取兩邊最大
func candy(ratings []int) int {
	need := make([]int, len(ratings))
	sum := 0

	for i := range ratings {
		need[i] = 1
	}

	// 前到後, 算右邊大的情況
	for i := range len(ratings) - 1 {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 後到前, 算左邊大的情況
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}

	for i := range ratings {
		sum += need[i]
	}

	return sum
}
