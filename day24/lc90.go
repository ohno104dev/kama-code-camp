package day24

import (
	"slices"
	"sort"
)

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	path2 []int
	res2  [][]int
	used  []bool
)

// 回溯
// 樹層剪枝(40)概念
func subsetsWithDup(nums []int) [][]int {
	path2, res2, used = make([]int, 0, len(nums)), make([][]int, 0, len(nums)), make([]bool, len(nums))
	sort.Ints(nums) // 排序剪枝
	backtracing(nums, 0)

	return res2
}

func backtracing(nums []int, start int) {
	res2 = append(res2, slices.Clone(path2))

	if start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		// used[i - 1] == true，說明同一樹枝candidates[i - 1]使用過
		// used[i - 1] == false，說明同一樹層candidates[i - 1]使用過 (剪枝)
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}

		path2 = append(path2, nums[i])
		used[i] = true
		backtracing(nums, i+1)
		path2 = path2[:len(path2)-1]

		used[i] = false
	}
}
