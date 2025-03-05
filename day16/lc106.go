package day16

// 遞迴解
// 注意分割區間
// 1. 確認函數的參數與返回值
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 2. 確認終止條件
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}

	// 3. 確認單層遞迴的邏輯
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]

	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}
