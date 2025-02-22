package day7

// Time Complexity: O(N^2+N^2)
// Space Complexity: O(N)

// 分散 Time Complexity
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dict := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			dict[a+b] += 1
		}
	}

	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			if v, ok := dict[0-(c+d)]; ok {
				count += v
			}
		}
	}

	return count
}
