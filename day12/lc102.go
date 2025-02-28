package day12

// BFS
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			size--

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}

	return ans
}
