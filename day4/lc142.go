package day4

// Time Complexity: O(N)
// Space Complexity: O(1)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// 快慢指針
// 數學證明
func detectCycle(head *ListNode2) *ListNode2 {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		// 找到環的相遇點
		if fast == slow {
			// 找環的起點
			for fast != head {
				fast = fast.Next
				head = head.Next
			}
			return fast
		}
	}

	return nil
}
