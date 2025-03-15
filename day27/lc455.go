package day27

import "sort"

// Time Complexity: O(N * long(N)) // 排序
// Space Complexity: O(1)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0

	p := len(g) - 1
	for c := len(s) - 1; c >= 0 && p >= 0; p-- {
		if s[c] >= g[p] {
			res++
			c--
		}

	}

	return res
}
