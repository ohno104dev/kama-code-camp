package day29

// Time Complexity: O(N)
// Space Complexity: O(1)

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for i := range bills {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five == 0 {
				return false
			}
			ten++
			five--
		} else {
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
			// 不會用20找零
		}
	}

	return true
}
