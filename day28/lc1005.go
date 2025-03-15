package day28

import (
	"math"
	"sort"
)

// Time Complexity: O(N)
// Space Complexity: O(1)

func largestSumAfterKNegations(nums []int, k int) int {

	// 如果只做一般排序, [-8, -1, 2] 負值取正後,原本正值最小不一定還會是最小
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := range nums {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 != 0 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	res := 0
	for _, v := range nums {
		res += v
	}

	return res
}
