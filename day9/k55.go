package day9

import "fmt"

// https://kamacoder.com/problempage.php?pid=1065
// Time Complexity: O(N)
// Space Complexity: O(N)

// 雙指針
// 多反轉解題思路
func rightReverseString() {
	var str string
	var target int

	fmt.Scanln(&target)
	fmt.Scanln(&str)

	strByte := []byte(str)

	reverse1(strByte, 0, len(strByte)-1)
	reverse1(strByte, 0, target-1)
	reverse1(strByte, target, len(strByte)-1)

	fmt.Printf("%s", string(strByte))
}

func reverse1(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
