package day16

// 注意分割區間
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}

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
