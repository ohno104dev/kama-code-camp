package day11

import "container/heap"

// Time Complexity: O(N*Log(K)), K小有優勢
// Space Complexity: O(K)

// 最小頂堆解 (保留前K大)
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)

	//所有元素入堆
	for key, val := range map_num {
		heap.Push(h, [2]int{key, val})
		// 把len(nums)-K小的都推出去,剩下前K大的
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := range k {
		// 返回堆中前K大的元素
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res
}

type IHeap [][2]int // [][key, value]

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// import "sort"

// // Time Complexity: O(N*Log(N))
// // Space Complexity: O(K)

// // 一般排序解
// func topKFrequent(nums []int, k int) []int {
// 	ans := []int{}
// 	map_num := map[int]int{}

// 	for _, item := range nums {
// 		map_num[item]++
// 	}

// 	for key, _ := range map_num {
// 		ans = append(ans, key)
// 	}

// 	sort.Slice(ans, func(a, b int) bool {
// 		return map_num[ans[a]] > map_num[ans[b]]
// 	})

// 	return ans[:k]
// }
