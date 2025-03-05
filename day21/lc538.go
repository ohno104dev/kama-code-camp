package day21

// 遞迴解
// 二叉搜索樹特性
// 反中序(右中左) => 大到小
// 使用指針(98)概念, 累加前後節點
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	traversal(root, &pre)
	return root
}

func traversal(node *TreeNode, pre *int) {
	if node == nil {
		return
	}

	traversal(node.Right, pre)
	node.Val += *pre
	*pre = node.Val
	traversal(node.Left, pre)
}
