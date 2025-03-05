package day18

import (
	"math"
)

// 遞迴解
// 二叉搜索樹中序特性解(从小到大有序)
// 使用指針比較前後節點的(98)概念
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	min := math.MaxInt64

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}

		prev = node
		traversal(node.Right)
	}

	traversal(root)
	return min
}
