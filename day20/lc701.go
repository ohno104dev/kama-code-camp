package day20

// 遞迴解
// 二叉搜索樹特性
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到要插入的點
	if root == nil {
		node := &TreeNode{Val: val}
		return node
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
