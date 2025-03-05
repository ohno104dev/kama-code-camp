package day17

// 遞迴解
// 注意分割區間
// 注意題目適合哪種序
// 1. 確認函數的參數與返回值
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 2. 確認終止條件
	if len(nums) == 0 {
		return nil
	}

	// 3. 確認單層遞迴的邏輯(此題適合前 序)
	idx := findMax(nums)
	root := &TreeNode{
		Val:   nums[idx],
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}

	return root
}

func findMax(nums []int) int {
	idx := 0
	for i, v := range nums {
		if nums[idx] < v {
			idx = i
		}
	}
	return idx
}
