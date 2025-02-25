package day10

type MyStack1 struct {
	queue []int
}

func Constructor1() MyStack1 {
	return MyStack1{
		queue: make([]int, 0),
	}
}

func (this *MyStack1) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack1) Pop() int {
	n := len(this.queue) - 1
	for n != 0 {
		val := this.queue[0]
		this.queue = this.queue[1:] // 移除
		this.queue = append(this.queue, val)
		n--
	}
	val := this.queue[0]
	this.queue = this.queue[1:] // 移除
	return val
}

func (this *MyStack1) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

func (this *MyStack1) Empty() bool {
	return len(this.queue) == 0
}
