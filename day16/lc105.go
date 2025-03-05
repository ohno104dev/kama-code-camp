package day16

// 遞迴解
// 注意分割區間
// 1. 確認函數的參數與返回值
func buildTree1(preorder []int, inorder []int) *TreeNode {
	// 2. 確認終止條件
	if len(preorder) == 0 {
		return nil
	}

	// 3. 確認單層遞迴的邏輯
	root := &TreeNode{Val: preorder[0]}
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree1(preorder[1:pos+1], inorder[:pos])
			root.Right = buildTree1(preorder[pos+1:], inorder[pos+1:])
		}
	}

	return root
}
