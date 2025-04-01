package day48

// Time Complexity: O(n)
// Space Complexity: O(n)

// 單調棧適合求當前元素左邊或右邊第一個比元素大或小的數值
// 遞增棧
// 使用(739)概念
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		if nums2[i] <= nums2[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && nums2[i] > nums2[top] {
				if v, ok := m[nums2[top]]; ok {
					res[v] = nums2[i]
				}
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}
