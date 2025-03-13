package day25

// Time Complexity: O(n!)
// Space Complexity: O(n)

// 回溯
// 二維遞迴
// 解數獨

func solveSudoku(board [][]byte) {
	var backtracking func(oard [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := range 9 {
			for j := range 9 {
				if board[i][j] != '.' {
					continue
				}

				// 嘗試1~9, 字符字面量的默認類型是rune
				for k := '1'; k <= '9'; k++ {
					if isValid2(i, j, byte(k), board) == true {
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false // 這格沒有數字符合
			}
		}

		return true // 填完所有空格
	}

	backtracking(board)
}

func isValid2(row, col int, k byte, board [][]byte) bool {
	// 橫
	for i := range 9 {
		if board[row][i] == k {
			return false
		}
	}

	// 直
	for i := range 9 {
		if board[i][col] == k {
			return false
		}
	}

	// 九宮格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3

	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}
