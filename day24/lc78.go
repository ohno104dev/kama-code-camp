package day24

import "slices"

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	res1  [][]int
	path1 []int
)

// 回溯
func subsets(nums []int) [][]int {
	path1, res1 = make([]int, 0), make([][]int, 0, len(nums))
	subdfs(nums, 0)

	return res1
}

func subdfs(nums []int, start int) {
	res1 = append(res1, slices.Clone(path1))
	if start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		path1 = append(path1, nums[i])
		subdfs(nums, i+1)
		path1 = path1[:len(path1)-1]
	}
}
