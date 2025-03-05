package day17

import (
	"math"
)

// 遞迴解
// 二叉搜索樹特性
func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

// 1. 確認函數的參數與返回值
func check(node *TreeNode, min, max int64) bool {
	// 2. 確認終止條件
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}

	// 3. 確認單層遞迴的邏輯
	return check(node.Left, min, int64(node.Val)) && check(node.Right, int64(node.Val), max)
}

// // 遞迴解
// // 二叉搜索樹中序特性解(从小到大有序)
// // 指針前後節點比較
// func isValidBST(root *TreeNode) bool {
// 	var prev *TreeNode // 紀錄上一個節點
// 	var travel func(node *TreeNode) bool
// 	travel = func(node *TreeNode) bool {
// 		if node == nil {
// 			return true
// 		}
// 		leftRes := travel(node.Left)
// 		// 當前值小於等於前一個節點的值，返回false
// 		if prev != nil && node.Val <= prev.Val {
// 			return false
// 		}
// 		prev = node
// 		rightRes := travel(node.Right)
// 		return leftRes && rightRes
// 	}
// 	return travel(root)
// }
