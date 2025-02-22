package day7

import "sort"

// Time Complexity: O(N^2)
// Space Complexity: O(1)

// 雙指針
// 去重重點
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// 剪枝
		if nums[i] > 0 {
			break
		}
		// 去重a
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for right > left {
			n1, n2, n3 := nums[i], nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重b
				for left < right && nums[left] == n2 {
					left++
				}
				// 去重c
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
