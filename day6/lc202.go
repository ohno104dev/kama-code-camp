package day6

// Time Complexity: O(Log(N)) // 每一步要處理的數字都會減少，導致所採取的步驟有一個上限
// Space Complexity: O(Log(N)) // 對於遞歸深度

func isHappy(n int) bool {
	m := make(map[int]bool) // 判斷是否循環

	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}

	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
