package day27

import "math"

// Time Complexity: O(N)
// Space Complexity: O(1)

// 只要求最大值, 不是最大子串
func maxSubArray(nums []int) int {
	max := math.MinInt32
	cur := 0

	for _, v := range nums {
		cur += v
		if cur > max {
			max = cur
		}

		// 當前為負數不會對後面的元素起到增大總和的作用, 重新計算
		if cur < 0 {
			cur = 0
		}
	}

	return max
}
