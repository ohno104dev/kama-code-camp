package day39

// Time Complexity: O(n)
// Space Complexity: O(n)

// dp[i]: 考慮下標i以內的房間, 最多偷竊金額為dp[i]
// 遞推公式: max(dp[i-1], dp[i-2]+nums[i])
// dp初始化: dp[0]=0
// 遍歷順序: 小到大

// 使用(198)概念
func rob2(nums []int) int {
	if len(nums) <= 1 {
		return rob1(nums)
	}

	//	把環攤平, 分開考慮
	res1 := rob1(nums[:len(nums)-1])
	res2 := rob1(nums[1:])
	return max(res1, res2)
}

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}
