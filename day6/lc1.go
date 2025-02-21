package day6

// Time Complexity: O(N)
// Space Complexity: O(N)

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for idx, val := range nums {
		if preIdx, ok := record[target-val]; ok {
			return []int{preIdx, idx}
		} else {
			record[val] = idx
		}
	}

	return []int{}
}
