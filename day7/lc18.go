package day7

import "sort"

// Time Complexity: O(N^3)
// Space Complexity: O(1)

// 雙指針
// 去重重點
// 使用三數之和(15)概念
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for a := 0; a < len(nums)-3; a++ {
		// 剪枝a, 考慮負數
		if target > 0 && nums[a] > 0 && nums[a] > target {
			break
		}
		// 去重a
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for b := a + 1; b < len(nums)-2; b++ {
			// 剪枝[a,b], 考慮負數
			if target > 0 && nums[a]+nums[b] > 0 && nums[a]+nums[b] > target {
				break
			}
			// 去重b
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			c := b + 1
			d := len(nums) - 1

			for c < d {
				n1, n2, n3, n4 := nums[a], nums[b], nums[c], nums[d]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					d--
				} else if sum < target {
					c++
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					// 去重c
					for c < d && n3 == nums[c+1] {
						c++
					}
					// 去重d
					for d > c && n4 == nums[d-1] {
						d--
					}
					// 移動指針
					c++
					d--
				}
			}
		}
	}
	return res
}
