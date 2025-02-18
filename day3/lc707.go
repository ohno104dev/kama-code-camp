package day3

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	dummy *Node
	Size  int // 額外判斷長度
}

func Constructor() MyLinkedList {
	dummy := &Node{} // 使用虛擬頭節點

	return MyLinkedList{
		dummy,
		0,
	}
}

// Time Complexity: O(N)
func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}

	cur := this.dummy.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

// Time Complexity: O(1)
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{}
	node.Val = val
	node.Next = this.dummy.Next
	this.dummy.Next = node
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummy
	for cur.Next != nil {
		cur = cur.Next
	}

	node := &Node{}
	node.Val = val
	cur.Next = node
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	node := &Node{}
	node.Val = val
	cur := this.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}

	cur := this.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
		this.Size--
	}
}
