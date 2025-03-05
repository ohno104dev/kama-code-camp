package day18

// 遞迴解
// 注意題目適合哪種序
// 1. 確認函數的參數與返回值
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 2. 確認終止條件
	if root == nil || root == p || root == q {
		return root
	}

	// 3. 確認單層遞迴的邏輯(此題適合後 序)
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}
