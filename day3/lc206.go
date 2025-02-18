package day3

// Time Complixity: O(N)
// Space Complixity: O(1)

type listNode2 struct {
	Val  int
	Next *listNode2
}

// 雙指針
func reverseList(head *listNode2) *listNode2 {
	var prev *listNode2 = nil
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
