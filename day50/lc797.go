package day50

import "slices"

// Time Complexity: O(2^n)
// Space Complexity: O(2^n)
// DFS
func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	dfs(0, graph, []int{}, &res)

	return res
}

func dfs(node int, graph [][]int, path []int, res *[][]int) {
	path = append(path, node)
	if node == len(graph)-1 {
		*res = append(*res, slices.Clone(path))
		return
	}

	for _, next := range graph[node] {
		dfs(next, graph, path, res)
	}
	path = path[:len(path)-1]
}
