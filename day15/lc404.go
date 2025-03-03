package day15

// // 遞迴解
// // 注意題目適合哪種序
// func sumOfLeftLeaves(root *TreeNode) int {
// 	return sums(root)
// }

// // 1. 確認函數的參數與返回值
// func sums(node *TreeNode) int {
// 	// 2. 確認終止條件
// 	if node == nil {
// 		return 0
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合後 序)
// 	lv := sums(node.Left)
// 	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
// 		lv = node.Left.Val
// 	}

// 	rv := sums(node.Right)

// 	return lv + rv //中
// }

// 迭代解
// 前序
func sumOfLeftLeaves(root *TreeNode) int {
	st := []*TreeNode{root}
	res := 0

	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}

		if node.Right != nil {
			st = append(st, node.Right)
		}

		if node.Left != nil {
			st = append(st, node.Left)
		}
	}

	return res
}
