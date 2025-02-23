package day8

import "fmt"

// https://kamacoder.com/problempage.php?pid=1064

func replaceNumber() {
	var strByte []byte

	fmt.Scanln(&strByte)

	for i := 0; i < len(strByte); i++ {
		if strByte[i] <= '9' && strByte[i] >= '0' {
			replace := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strByte = append(strByte[:i], append(replace, strByte[i+1:]...)...)

			i = i + len(replace) - 1
		}
	}

	fmt.Printf("%s", string(strByte))
}
