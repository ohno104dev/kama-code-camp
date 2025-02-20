package day4

// Time Complexity: O(A＋B)
// Space Complexity: O(1)

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

// 使用快慢指針(19)對齊概念
func getIntersectionNode(headA, headB *ListNode3) *ListNode3 {
	curA := headA
	curB := headB

	lenA, lenB := 0, 0

	for curA != nil {
		curA = curA.Next
		lenA++
	}

	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode3
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// 鏈表拼接
// https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/12624/intersection-of-two-linked-lists-shuang-zhi-zhen-l/
// func getIntersectionNode(headA, headB *ListNode3) *ListNode3 {
//     l1,l2 := headA, headB
//     for l1 != l2 {
//         if l1 != nil {
//             l1 = l1.Next
//         } else {
//             l1 = headB
//         }

//         if l2 != nil {
//             l2 = l2.Next
//         } else {
//             l2 = headA
//         }
//     }

//     return l1
// }
