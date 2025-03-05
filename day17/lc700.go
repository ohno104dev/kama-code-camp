package day17

// // 遞迴解
// // 二叉搜索樹特性
// func searchBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil || root.Val == val {
// 		return root
// 	}

// 	if root.Val > val {
// 		return searchBST(root.Left, val)
// 	}

// 	return searchBST(root.Right, val)
// }

// 迭代解
// 二叉搜索樹特性
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}

	return root
}
