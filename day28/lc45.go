package day28

// Time Complexity: O(N)
// Space Complexity: O(1)

// 轉換思路, 只管能不能cover到最後一個位置
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	lastDist := 0
	curDist := 0
	step := 0

	for i := range nums {
		// 超過上一次覆蓋範圍還沒到最後
		if i == lastDist+1 {
			step++
			lastDist = curDist
		}

		curDist = max(nums[i]+i, curDist)
	}

	return step
}
