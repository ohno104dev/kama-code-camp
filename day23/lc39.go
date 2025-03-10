package day23

import (
	"slices"
)

// Time Complexity: O(n * 2^n)
// Space Complexity: O(target)

var (
	path []int
	res  [][]int
)

// 回溯
func combinationSum(candidates []int, target int) [][]int {
	path, res = make([]int, 0), make([][]int, 0, len(candidates))
	sumdfs(&candidates, 0, target, 0)

	return res
}

func sumdfs(candidates *[]int, sum int, target int, start int) {
	if sum > target {
		return
	}

	if sum == target {
		res = append(res, slices.Clone(path))
		return
	}

	for i := start; i < len(*candidates); i++ {
		path = append(path, (*candidates)[i])
		sum += (*candidates)[i]
		sumdfs(candidates, sum, target, i)
		sum -= (*candidates)[i]
		path = path[:len(path)-1]
	}
}

// // 排序剪枝
// func combinationSum(candidates []int, target int) [][]int {
// 	path, res = make([]int, 0), make([][]int, 0, len(candidates))
// 	sort.Ints(candidates) // 排序剪枝
// 	sumdfs(candidates, 0, target)

// 	return res
// }

// func sumdfs(candidates []int, start int, target int) {
// 	if target == 0 {
// 		res = append(res, slices.Clone(path))
// 		return
// 	}

// 	for i := start; i < len(candidates); i++ {
// 		if candidates[i] > target {
// 			return
// 		}

// 		path = append(path, candidates[i])
// 		sumdfs(candidates, i, target-candidates[i])
// 		path = path[:len(path)-1]
// 	}
// }
