package day30

import "sort"

// Time Complexity: O(N*Log(N))
// Space Complexity: O(N)

func findMinArrowShots(points [][]int) int {
	res := 1

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return res
}
