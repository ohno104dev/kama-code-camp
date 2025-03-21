package day34

// Time Complexity: O(n^2)
// Space Complexity: O(n)

// 確認dp數組含義: 為節點組成的二叉搜索數的個數
// 遞推公式: dp[i]+=dp[j-1]*dp[i-j]
// dp數組如何初始化: dp[0]=1, dp[1]=1, dp[2]=1
// 遍歷順序: 從小到大
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//一共i個節點，對於根節點j, 左子樹的節點為j-1，右子樹的節點為i-j
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
