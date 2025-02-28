package day4

// Time Complexity: O(N)
// Space Complexity: O(1)

// 快慢指針
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy

	// fast多走一步, slow的next會是第n個node
	move := n + 1
	for move > 0 && fast != nil {
		fast = fast.Next
		move--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
