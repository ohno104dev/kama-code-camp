package day31

import "sort"

// Time Complexity: O(N*Log(N))
// Space Complexity: O(Log(N))

// 使用(452)概念
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0]) // 放入第一個區間

	for i := 1; i < len(intervals); i++ {
		// 每次檢查最後的區間
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
