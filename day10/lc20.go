package day10

// Time Complexity: O(N)
// Space Complexity: O(N)

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r) // 左括號轉成右括號簡化比對
		} else {
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
