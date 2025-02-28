package day13

// // 遞迴解
// // 注意題目適合哪種序
// func invertTree(root *TreeNode) *TreeNode {
// 	invertversal(root)
// 	return root
// }

// // 1. 確認函數的參數與返回值
// func invertversal(cur *TreeNode) {
// 	// 2. 確認終止條件
// 	if cur == nil {
// 		return
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合前/後 序)
// 	cur.Left, cur.Right = cur.Right, cur.Left
// 	invertversal(cur.Left)
// 	invertversal(cur.Right)
// }

// BFS解
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			size--
		}
	}

	return root
}
