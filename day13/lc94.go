package day13

// 遞迴三步
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}
	inversal(root, &ans)
	return ans
}

// 1. 確認遞迴的參數與返回值
func inversal(cur *TreeNode, ans *[]int) {
	// 2. 確認終止條件
	if cur == nil {
		return
	}

	// 3. 確認單層遞迴的邏輯
	inversal(cur.Left, ans)
	*ans = append(*ans, cur.Val)
	inversal(cur.Right, ans)
}
