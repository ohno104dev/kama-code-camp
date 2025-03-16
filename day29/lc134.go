package day29

// Time Complexity: O(N)
// Space Complexity: O(1)

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	cur := 0
	total := 0
	for i := range gas {
		cur += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
