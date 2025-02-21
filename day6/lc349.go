package day6

// Time Complexity: O(N)
// Space Complexity: O(N)

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 0, max(len(nums1), len(nums2)))

	for _, r := range nums1 {
		set[r] = struct{}{}
	}

	for _, r := range nums2 {
		if _, ok := set[r]; ok {
			res = append(res, r)
			delete(set, r) // 去重
		}
	}

	return res
}
