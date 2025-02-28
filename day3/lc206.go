package day3

// Time Complixity: O(N)
// Space Complixity: O(1)

// 雙指針
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// func reverseList(head *ListNode) (prev *ListNode) {
// 	for head != nil {
// 		head.Next, prev, head = prev, head, head.Next
// 	}
// 	return
// }
