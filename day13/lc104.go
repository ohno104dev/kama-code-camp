package day13

// // 遞迴解
// // 注意題目適合哪種序
// func maxDepth(root *TreeNode) int {
// 	return depth(root)
// }

// // 1. 確認函數的參數與返回值
// func depth(node *TreeNode) int {
// 	// 2. 確認終止條件
// 	if node == nil {
// 		return 0
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合前/後 序)
// 	l := depth(node.Left)
// 	r := depth(node.Right)
// 	return 1 + max(l, r)
// }

// BFS解
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		depth++
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
	}

	return depth
}
