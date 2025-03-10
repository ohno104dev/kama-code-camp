package day23

import (
	"slices"
	"sort"
)

// Time Complexity: O(n * 2^n)
// Space Complexity: O(n)

var (
	res1  [][]int
	path1 []int
	used  []bool //樹層剪枝
)

// 回溯
// 樹層剪枝概念
func combinationSum2(candidates []int, target int) [][]int {
	res1, path1 = make([][]int, 0), make([]int, 0, len(candidates))
	used = make([]bool, len(candidates))
	sort.Ints(candidates) // 排序剪枝
	sum2dfs(&candidates, 0, 0, target)

	return res1
}

func sum2dfs(candidates *[]int, start int, sum int, target int) {
	if sum == target {
		res1 = append(res1, slices.Clone(path1))

		return
	}

	for i := start; i < len(*candidates); i++ {
		if sum+(*candidates)[i] > target {
			return
		}

		// 樹層剪枝
		// used[i-1] == false 才是新的樹層分支
		if i > 0 && (*candidates)[i] == (*candidates)[i-1] && used[i-1] == false {
			continue
		}

		// 一般處理
		// if i != start && (*candidates)[i] == (*candidates)[i-1] {
		//     continue
		// }

		path1 = append(path1, (*candidates)[i])
		used[i] = true
		sum += (*candidates)[i]
		sum2dfs(candidates, i+1, sum, target)
		used[i] = false
		sum -= (*candidates)[i]
		path1 = path1[:len(path1)-1]
	}
}
