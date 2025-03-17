package day30

import "sort"

// Time Complexity: O(N*Log(N))
// Space Complexity: O(N)

// 使用二分查找(452)概念
func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			res++
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}

	return res
}
