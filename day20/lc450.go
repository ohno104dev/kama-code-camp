package day20

// 遞迴解
// 二叉搜索樹特性
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		//處理邏輯在終止條件
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Right != nil && root.Left == nil {
			return root.Right
		} else {
			// 把整個子樹接過去
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
