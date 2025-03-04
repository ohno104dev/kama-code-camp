package day16

// 遞迴解
// 回溯概念
// 注意題目適合哪種序
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	findPath(root, &res, []int(nil), targetSum)

	return res
}

// 1. 確認函數的參數與返回值
func findPath(node *TreeNode, res *[][]int, curPath []int, remaining int) {
	// 2. 確認終止條件
	if node == nil {
		return
	}

	// 3. 確認單層遞迴的邏輯(此題適合前 序)
	remaining -= node.Val
	curPath = append(curPath, node.Val)
	if node.Left == nil && node.Right == nil && remaining == 0 {
		*res = append(*res, append([]int{}, curPath...))
		curPath = curPath[:len(curPath)-1]
	}
	findPath(node.Left, res, curPath, remaining)
	findPath(node.Right, res, curPath, remaining)
}
