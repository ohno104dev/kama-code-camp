package day4

// Time Complexity: O(N)
// Space Complexity: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意終止條件, 交換順序
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp
		tmp.Next = tmp1

		cur = cur.Next.Next
	}

	return dummy.Next
}
