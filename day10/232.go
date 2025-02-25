package day10

type MyStack []int

func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return val
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Size() int {
	return len(*s)
}

func (s *MyStack) Empty() bool {
	return s.Size() == 0
}

type MyQueue struct {
	in  *MyStack
	out *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  &MyStack{},
		out: &MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.out.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.out.Peek()
}

// 複用寫法
// func (this *MyQueue) Peek() int {
//     val := this.Pop()
//     this.out.Push(val)
//     return val
// }

func (this *MyQueue) Empty() bool {
	return this.out.Empty() && this.in.Empty()
}

// 填充輸出Stack
func (this *MyQueue) fillStackOut() {
	if this.out.Empty() {
		for !this.in.Empty() {
			val := this.in.Pop()
			this.out.Push(val)
		}
	}
}
