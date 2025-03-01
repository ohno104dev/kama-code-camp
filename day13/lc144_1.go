package day13

// 迭代作法
// 利用stack
func preorderTraversal1(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)

		// 取出的順序左右會相反
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}
