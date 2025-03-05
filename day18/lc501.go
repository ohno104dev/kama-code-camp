package day18

// 遞迴解
// 注意題目適合哪種序
// 計數法, 不使用額外空間, 只需要遍歷一次
// 使用指針比較前後節點的(98)概念
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)

		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}

			max = count
		}
		prev = node
		traversal(node.Right)
	}

	traversal(root)
	return res
}
