package day16

// // 遞迴解
// // 注意題目適合哪種序
// var depth int
// var res int

// func findBottomLeftValue(root *TreeNode) int {
// 	depth, res = 0, 0
// 	dfs(root, 1)

// 	return res
// }

// // 1. 確認函數的參數與返回值
// func dfs(node *TreeNode, dep int) {
// 	// 2. 確認終止條件
// 	if node == nil {
// 		return
// 	}

// 	if node.Left == nil && node.Right == nil && depth < dep {
// 		depth = dep
// 		res = node.Val
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合前中後 序)(都是左邊開始)
// 	dfs(node.Left, dep+1)
// 	dfs(node.Right, dep+1)
// }

// BFS解
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		level := len(queue)
		for i := range level {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
