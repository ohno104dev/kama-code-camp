package day49

// Time Complexity: O(n)
// Space Complexity: O(n)

// 單調棧適合求當前元素左邊或右邊第一個比元素大或小的數值
// 遞減棧

func largestRectangleArea(heights []int) int {
	res := 0
	heights = append([]int{0}, heights...) // 避免數組本身為降序的邊界問題
	heights = append(heights, 0)           // 避免數組本身為升序的邊界問題
	st := []int{0}

	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[st[len(st)-1]] {
			st = append(st, i)
		} else if heights[i] == heights[st[len(st)-1]] {
			st = st[:len(st)-1]
			st = append(st, i)
		} else {
			for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
				mid := st[len(st)-1]
				st = st[:len(st)-1]
				if len(st) > 0 {
					left := st[len(st)-1]
					right := i
					w := right - left - 1
					h := heights[mid]
					res = max(res, w*h)
				}
			}
		}
		st = append(st, i)
	}

	return res
}

// 雙指針
// func largestRectangleArea(heights []int) int {
// 	size := len(heights)
// 	minLeftIndex := make([]int, size)
// 	minRightIndex := make([]int, size)

// 	// 紀錄每個柱子, 左邊第一個小於該柱子的index
// 	minLeftIndex[0] = -1 // 注意初始化, 防止while死循環
// 	for i := 1; i < size; i++ {
// 		t := i - 1
// 		// 向左尋找
// 		for t >= 0 && heights[t] >= heights[i] {
// 			t = minLeftIndex[t]
// 		}
// 		minLeftIndex[i] = t
// 	}
// 	// 紀錄每個柱子, 右邊第一個小於該柱子的index
// 	minRightIndex[size-1] = size // 注意初始化, 防止while死循環
// 	for i := size - 2; i >= 0; i-- {
// 		t := i + 1
// 		// 向右尋找
// 		for t < size && heights[t] >= heights[i] {
// 			t = minRightIndex[t]
// 		}
// 		minRightIndex[i] = t
// 	}

// 	result := 0
// 	for i := 0; i < size; i++ {
// 		sum := heights[i] * (minRightIndex[i] - minLeftIndex[i] - 1)
// 		result = max(sum, result)
// 	}
// 	return result
// }
