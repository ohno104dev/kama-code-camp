package day13

// 迭代作法
// 利用stack
// 翻轉解
func postorderTraversal1(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)

		// 1. 中右左
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(ans) // 2. 中右左 => 左右中
	return ans
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
