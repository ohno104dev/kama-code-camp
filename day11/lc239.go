package day11

// Time Complexity: O(N)
// Space Complexity: O(K)

// 單調隊列
func maxSlidingWindow(nums []int, k int) []int {
	d, max := &deque{}, []int{}
	for i, n := range nums {
		d.push(n)
		if i-k >= -1 { // index跟長度差
			max = append(max, d.front())
			d.pop(nums[i-k+1])
		}
	}

	return max
}

type deque []int

func (d *deque) front() int { return (*d)[0] }
func (d *deque) back() int  { return (*d)[len(*d)-1] }
func (d *deque) push(n int) {
	for len(*d) > 0 && n > d.back() {
		*d = (*d)[:len(*d)-1]
	}
	*d = append(*d, n)
}
func (d *deque) pop(n int) {
	if len(*d) > 0 && n == d.front() {
		*d = (*d)[1:]
	}
}
