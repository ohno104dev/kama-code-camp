package day29

import (
	"slices"
	"sort"
)

// Time Complexity: O(N*log(N)+ N^2)
// Space Complexity: O(N)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// insert是N^2
	// for i, p := range people {
	// 	copy(people[p[1]+1:i+1], people[p[1]:i+1])
	// 	people[p[1]] = p
	// }
	// return people

	res := make([][]int, 0, len(people))
	for _, p := range people {
		res = slices.Insert(res, p[1], p) // 直接插入
	}

	return res
}
