package day3

// Time Complexity: O(N)
// Space Complexity: O(1)

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 前一個Node檢查下一個的作法比判斷Node本身操作簡潔
// 使用dummy head
func removeElements(head *ListNode1, val int) *ListNode1 {
	//注意是指針宣告
	dummy := &ListNode1{}
	dummy.Next = head

	curr := dummy // 前一個Node檢查下一個的作法比判斷Node本身操作簡潔
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

// func removeElements(head *ListNode, val int) *ListNode {
// 	// 處理head是val的情況
// 	for head != nil && head.Val == val {
// 		head = head.Next
// 	}
// 	curr := head
// 	for curr != nil && curr.Next != nil {
// 		if curr.Next.Val == val {
// 			curr.Next = curr.Next.Next
// 		} else {
// 			curr = curr.Next
// 		}
// 	}
// 	return head
// }
