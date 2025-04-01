package day48

// Time Complexity: O(n)
// Space Complexity: O(n)

// 單調棧適合求當前元素左邊或右邊第一個比元素大或小的數值
// 遞增棧
// 使用(496)概念
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

	stack := []int{0}
	for i := range len(nums) * 2 {
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = nums[i%len(nums)]
		}

		stack = append(stack, i%len(nums))
	}

	return res
}
