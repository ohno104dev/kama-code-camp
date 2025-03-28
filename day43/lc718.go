package day43

// Time Complexity: O(n*m)
// Space Complexity: O(n*m)

// dp[i][j]: 以nums1[i-1]為結尾和nums2[j-1]為結尾的最長重複子數組的長度 (可以避免初始化dp[0][j]與dp[i][0])
// 遞推公式: dp[i][j]=dp[i-1][j-1]+1
// dp初始化: dp[:] = 0
// 遍歷順序: 前到後, 上到下

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
