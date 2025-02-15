package day2

// Time Complixity: O(N^2)
// Space Complixity: O(N^2)

// 邊界處理, 循環不變量(統一處理規則)
func generateMatrix(n int) [][]int {
	// 注意二维slice初始化寫法
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	sx, sy := 0, 0 // 記錄loop起始位置
	offset := 1    // 保持左閉右開
	num := 1       // 計算數值
	loop := n / 2

	for loop > 0 {
		x, y := sx, sy
		for x < n-offset {
			res[y][x] = num
			x++
			num++
		}

		for y < n-offset {
			res[y][x] = num
			y++
			num++
		}

		for x > sx {
			res[y][x] = num
			x--
			num++
		}

		for y > sy {
			res[y][x] = num
			y--
			num++
		}

		sx++
		sy++
		offset++
		loop--
	}

	// 處理奇數array中心點
	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}

	return res

}
