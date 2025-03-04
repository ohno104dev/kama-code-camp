package day16

// // 遞迴解
// // 回溯概念
// // 注意題目適合哪種序
func hasPathSum(root *TreeNode, targetSum int) bool {
	return traversal(root, targetSum)
}

// 1. 確認函數的參數與返回值
func traversal(node *TreeNode, remaining int) bool {
	// 2. 確認終止條件
	if node == nil {
		return false
	}

	// 3. 確認單層遞迴的邏輯(此題適合前 序)
	remaining -= node.Val
	if node.Left == nil && node.Right == nil {
		return remaining == 0
	}

	if traversal(node.Left, remaining) {
		return true
	}

	if traversal(node.Right, remaining) {
		return true
	}

	return false
}
