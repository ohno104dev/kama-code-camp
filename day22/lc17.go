package day22

// Time Complexity: O(3^m * 4^n) // m 是對應三個字母的數字個數,n 是對應四個字母的數字個數
// Space Complexity: O(3^m * 4^n)

var (
	path2 []byte
	res2  []string
	rec   []string
)

// 回溯
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	rec = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path2, res2 = make([]byte, 0), make([]string, 0)
	backtracking(digits, 0)

	return res2
}

func backtracking(digits string, index int) {
	if index == len(digits) {
		res2 = append(res2, string(path2))
		return
	}

	str := rec[int(digits[index]-'0')]
	for j := range len(str) {
		path2 = append(path2, str[j])
		backtracking(digits, index+1)
		path2 = path2[:len(path2)-1]
	}
}
