package day10

// Time Complexity: O(N)
// Space Complexity: O(N)

func removeDuplicates(s string) string {
	stack := make([]rune, 0)

	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}
