package day14

// // 遞迴解
// // 注意題目適合哪種序
// func minDepth(root *TreeNode) int {
// 	return mindepth(root)
// }

// // 1. 確認函數的參數與返回值
// func mindepth(node *TreeNode) int {
// 	// 2. 確認終止條件
// 	if node == nil {
// 		return 0
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合前/後 序)
// 	l := mindepth(node.Left)
// 	r := mindepth(node.Right)
// 	if node.Left == nil && node.Right != nil {
// 		return 1 + r
// 	}
// 	if node.Right == nil && node.Left != nil {
// 		return 1 + l
// 	}

// 	return 1 + min(l, r)
// }

// BFS解
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

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
