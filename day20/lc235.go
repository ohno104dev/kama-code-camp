package day20

// 遞迴解
// 二叉搜索樹特性
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}

	return root
}

// // 迭代解
// // 二叉搜索樹特性
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

//     for root != nil {
//         if root.Val > p.Val && root.Val > q.Val {
//             root = root.Left
//         } else if root.Val < p.Val && root.Val < q.Val {
//             root = root.Right
//         } else {
//             return root
//         }
//     }

//     return nil
// }
