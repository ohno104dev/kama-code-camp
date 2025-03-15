package day27

// Time Complexity: O(N)
// Space Complexity: O(1)

// case 1: 上下坡中有平台 (1-2-2-2-1)
// case 2: 數組首尾兩端 (2-2-5)
// case 3: 單調坡中有平台 (1-2-2-2-3-4)
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	ans := 1 // case2: 尾段必為擺動

	// case2: 計算頭段(可能是平的)
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}

	// case1: 計算中段
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++

			// case3: 只有在發生變化時更新
			prevDiff = diff
		}
	}

	return ans
}
