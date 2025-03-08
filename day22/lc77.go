package day22

import "slices"

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	path []int
	res  [][]int
)

// 回溯
func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	combdfs(n, k, 1)

	return res
}

func combdfs(n int, k int, start int) {
	if len(path) == k {
		res = append(res, slices.Clone(path)) // go 1.21

		return
	}

	// 剪枝
	// 1. 剩餘不足k
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		combdfs(n, k, i+1)
		path = path[:len(path)-1] // 回溯
	}
}
