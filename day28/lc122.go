package day28

// Time Complexity: O(N)
// Space Complexity: O(1)

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}

	return profit
}
