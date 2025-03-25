package day43

// Time Complexity: O(n)
// Space Complexity: O(1), 用dp數組: O(n)

// 使用(300)概念
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, count := 1, 1
	for i := range len(nums) - 1 {
		if nums[i+1] > nums[i] {
			count++
		} else {
			count = 1
		}
		if count > res {
			res = count
		}
	}

	return res
}
