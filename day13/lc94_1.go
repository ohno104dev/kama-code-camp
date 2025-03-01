package day13

// 迭代作法
// 利用stack
func inorderTraversal1(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	stack := []*TreeNode{} // 紀錄處理的節點
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}
