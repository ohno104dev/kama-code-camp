package day25

// Time Complexity: O(n!)
// Space Complexity: O(n)

// 回溯
// 三維陣列
// N皇后
func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := range n {
		chessboard[i] = make([]string, n)
		for j := range n {
			chessboard[i][j] = "."
		}
	}

	var backtracking func(int)
	backtracking = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := range n {
				tmp[i] = ""
				for j := range n {
					tmp[i] += chessboard[i][j]
				}
			}
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtracking(row + 1)
				chessboard[row][i] = "."
			}
		}
	}

	backtracking(0)
	return res
}

// 往前檢查
func isValid(n, row, col int, chessboard [][]string) bool {
	// 按行放置皇后，每一行只會放一個皇后，因此不會有行衝突

	// 檢查同一列(直)衝突
	for i := range row {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	// 檢查左上角衝突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	// 檢查右上角衝突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}
