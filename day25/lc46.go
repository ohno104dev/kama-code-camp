package day25

import "slices"

// Time Complexity: O(n!)
// Space Complexity: O(n)

var (
	res1  [][]int
	path1 []int
)

// 回溯
// 排列
func permute(nums []int) [][]int {
	res1, path1 = make([][]int, 0), make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	permutedfs(&nums, &used)

	return res1
}

func permutedfs(nums *[]int, used *[]bool) {
	if len(path1) == len(*nums) {
		res1 = append(res1, slices.Clone(path1))
		return
	}

	// 每次都從頭
	for i := range *nums {
		if (*used)[i] == true {
			continue
		}

		path1 = append(path1, (*nums)[i])
		(*used)[i] = true
		permutedfs(nums, used)
		path1 = path1[:len(path1)-1]
		(*used)[i] = false
	}
}
