package day15

// 遞迴解
// 注意題目適合哪種序
func isBalanced(root *TreeNode) bool {
	return balanced(root) > -1
}

// 1. 確認函數的參數與返回值
func balanced(node *TreeNode) int {
	// 2. 確認終止條件
	if node == nil {
		return 0
	}

	// 3. 確認單層遞迴的邏輯(此題適合後 序)
	l := balanced(node.Left)
	if l == -1 {
		return -1
	}

	r := balanced(node.Right)
	if r == -1 {
		return -1
	}

	// 高度差大於一不符合條件
	if l-r > 1 || r-l > 1 {
		return -1
	}

	return max(l, r) + 1
}
