package day25

import "slices"

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	path []int
	res  [][]int
)

// 回溯
// 樹層剪枝(40)概念變形
func findSubsequences(nums []int) [][]int {
	path, res = make([]int, 0), make([][]int, 0)
	subseqdfs(&nums, 0)
	return res
}

func subseqdfs(nums *[]int, start int) {
	if len(path) > 1 {
		res = append(res, slices.Clone(path))
	}

	used := make(map[int]bool, len(*nums))

	for i := start; i < len(*nums); i++ {
		if used[(*nums)[i]] {
			continue
		}

		if len(path) == 0 || (*nums)[i] >= path[len(path)-1] {
			path = append(path, (*nums)[i])
			used[(*nums)[i]] = true
			subseqdfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
}
