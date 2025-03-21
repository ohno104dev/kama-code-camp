package day32

// Time Complexity: O(N)
// Space Complexity: O(N)

// 確認dp數組含義: 第i個fib的數值為dp[i]
// 遞推公式: dp[i] = dp[i-1]+dp[i-2]
// dp數組如何初始化: dp[0]=1, dp[1]=1
// 遍歷順序: 前向後
func fib(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n, n)
	dp[0], dp[1] = 1, 1

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// Time Complexity: O(N)
// Space Complexity: O(1)
// func fib(n int) int {
//     if n < 2 {
//         return n
//     }
//     a, b, c := 0, 1, 0
//     for i := 1; i < n; i++ {
//         c = a + b
//         a, b = b, c
//     }
//         return c
// }
