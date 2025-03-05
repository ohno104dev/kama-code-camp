package day21

// 遞迴解
// 二叉搜索樹特性
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{Val: nums[idx]}

	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])

	return root
}
