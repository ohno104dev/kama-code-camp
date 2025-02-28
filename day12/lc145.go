package day12

// 遞迴三步
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}
	postversal(root, &ans)
	return ans
}

// 1. 確認遞迴函數的參數與返回值
func postversal(cur *TreeNode, ans *[]int) {
	// 2. 確認終止條件
	if cur == nil {
		return
	}

	// 3. 確認單層遞迴的邏輯
	postversal(cur.Left, ans)
	postversal(cur.Right, ans)
	*ans = append(*ans, cur.Val)
}
