package day13

// 遞迴三步
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int // ans ptr是nil, 或是用ans := []int{}
	preversal(root, &ans)

	return ans
}

// 1. 確認遞迴函數的參數與返回值
func preversal(cur *TreeNode, ans *[]int) {
	// 2. 確認終止條件
	if cur == nil {
		return
	}
	// 3.確認單層遞迴的邏輯
	*ans = append(*ans, cur.Val)
	preversal(cur.Left, ans)
	preversal(cur.Right, ans)
}
