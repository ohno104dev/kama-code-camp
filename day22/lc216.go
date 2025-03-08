package day22

import "slices"

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	path1 []int
	res1  [][]int
)

// 回溯
func combinationSum3(k int, n int) [][]int {
	res1, path1 = make([][]int, 0), make([]int, 0, k)
	sum3dfs(k, n, 1, 0)
	return res1
}

func sum3dfs(k, n, start, sum int) {
	if len(path1) == k {
		if sum == n {
			res1 = append(res1, slices.Clone(path1)) // go 1.21
		}
		return
	}

	for i := start; i <= 9; i++ {
		// 剪枝
		// 1. sum值超過
		// 2. 剩餘不足k
		if sum+i > n || 9-i+1 < k-len(path1) {
			break
		}
		path1 = append(path1, i)
		sum3dfs(k, n, i+1, sum+i)
		path1 = path1[:len(path1)-1] // 回溯
	}
}
