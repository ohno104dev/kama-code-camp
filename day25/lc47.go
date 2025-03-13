package day25

import (
	"slices"
	"sort"
)

// Time Complexity: O(n! * n)
// Space Complexity: O(n)

var (
	res2  [][]int
	path2 []int
)

// 回溯
// 排列
// 樹層剪枝(40)概念
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 排序剪枝
	res2, path2 = make([][]int, 0), make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	permute2dfs(&nums, &used)

	return res2
}

func permute2dfs(nums *[]int, used *[]bool) {
	if len(path2) == len(*nums) {
		res2 = append(res2, slices.Clone(path2))
		return
	}

	// 每次都從頭
	for i := 0; i < len(*nums); i++ {
		if i > 0 && (*nums)[i-1] == (*nums)[i] && (*used)[i-1] == false {
			continue
		}

		if (*used)[i] == true {
			continue
		}

		path2 = append(path2, (*nums)[i])
		(*used)[i] = true
		permute2dfs(nums, used)
		path2 = path2[:len(path2)-1]
		(*used)[i] = false
	}
}
