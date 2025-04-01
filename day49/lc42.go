package day49

// Time Complexity: O(n)
// Space Complexity: O(n)

// 單調棧適合求當前元素左邊或右邊第一個比元素大或小的數值
// 遞增棧
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	st := make([]int, len(height))
	var res int

	for i := 1; i < len(height); i++ {
		if height[i] < height[st[len(st)-1]] {
			st = append(st, i)
		} else if height[i] == height[st[len(st)-1]] {
			st = st[:len(st)-1]
			st = append(st, i)
		} else {
			for len(st) != 0 && height[i] > height[st[len(st)-1]] {
				top := st[len(st)-1]
				st = st[:len(st)-1]
				if len(st) != 0 {
					res += (min(height[i], height[st[len(st)-1]]) - height[top]) * (i - st[len(st)-1] - 1)
				}
			}

			st = append(st, i)
		}
	}

	return res
}

// 雙指針
// func trap(height []int) int {
// 	sum := 0
// 	n := len(height)
// 	lh := make([]int, n)
// 	rh := make([]int, n)
// 	lh[0] = height[0]
// 	rh[n-1] = height[n-1]
// 	for i := 1; i < n; i++ {
// 		lh[i] = max(lh[i-1], height[i])
// 	}
// 	for i := n - 2; i >= 0; i-- {
// 		rh[i] = max(rh[i+1], height[i])
// 	}
// 	for i := 1; i < n-1; i++ {
// 		h := min(rh[i], lh[i]) - height[i]
// 		if h > 0 {
// 			sum += h
// 		}
// 	}
// 	return sum
// }
