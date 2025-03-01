package day14

// // 遞迴解
// // 注意題目適合哪種序
// func isSymmetric(root *TreeNode) bool {
// 	return compare(root.Left, root.Right)
// }

// // 1. 確認函數的參數與返回值
// func compare(left, right *TreeNode) bool {
// 	// 2. 確認終止條件
// 	if left == nil && right != nil {
// 		return false
// 	}
// 	if left != nil && right == nil {
// 		return false
// 	}
// 	if left == nil && right == nil {
// 		return true
// 	}
// 	if left.Val != right.Val {
// 		return false
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合後序)
// 	outside := compare(left.Left, right.Right)
// 	inside := compare(left.Right, right.Left)
// 	return outside && inside
// }

// 迭代解
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}
