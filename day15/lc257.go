package day15

import "strconv"

// // 遞迴解
// // 回溯概念
// func binaryTreePaths(root *TreeNode) []string {
// 	ans := []string{}

// 	paths(root, "", &ans)

// 	return ans
// }

// // 1. 確認函數的參數與返回值
// func paths(cur *TreeNode, pre string, ans *[]string) {
// 	// 2. 確認終止條件
// 	if cur.Left == nil && cur.Right == nil {
// 		path := pre + strconv.Itoa(cur.Val)
// 		*ans = append(*ans, path)
// 		return
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合前 序)
// 	pre = pre + strconv.Itoa(cur.Val) + "->"
// 	if cur.Left != nil {
// 		paths(cur.Left, pre, ans)
// 	}

// 	if cur.Right != nil {
// 		paths(cur.Right, pre, ans)
// 	}
// }

// 迭代解
func binaryTreePaths(root *TreeNode) []string {
	st := []*TreeNode{root}
	paths := []string{""}
	res := []string{}

	for len(st) > 0 {
		node := st[len(st)-1]
		path := paths[len(paths)-1]
		st = st[:len(st)-1]
		paths = paths[:len(paths)-1]

		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}

		if node.Right != nil {
			st = append(st, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}

		if node.Left != nil {
			st = append(st, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}

	return res
}
