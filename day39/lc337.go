package day39

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n)
// Space Complexity: O(Log(n))

// []dp{}: 返回數組dp[0]:偷最大值, dp[1]:不偷最大值
// 遞推公式:
// 偷父節點: cur.Val + left[0] + right[0]
// 不偷父節點: slices.Max(left) + slices.Max(right)
// dp初始化: []dp{0,0}
// 遍歷順序: 小到大

// 使用(198)概念
// dp + Tree
func rob3(root *TreeNode) int {
	res := robtraversal(root)
	return slices.Max(res)
}

func robtraversal(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}

	left := robtraversal(cur.Left)
	right := robtraversal(cur.Right)

	rob := cur.Val + left[0] + right[0]            // 偷父節點
	notRob := slices.Max(left) + slices.Max(right) // 不偷父節點

	return []int{notRob, rob}
}
