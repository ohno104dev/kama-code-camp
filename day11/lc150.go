package day11

import "strconv"

// Time Complexity: O(N)
// Space Complexity: O(N)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {

		if val, err := strconv.Atoi(token); err == nil {
			stack = append(stack, val)
		} else {
			// err != nil 表示不是數字
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1] // 順序顛倒
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)

			case "-":
				stack = append(stack, num1-num2)

			case "*":
				stack = append(stack, num1*num2)

			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	return stack[0]
}
