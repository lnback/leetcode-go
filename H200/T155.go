package H200


type MinStack struct {
	stack  []Node
}
type Node struct{
	num int
	minNum int
}
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	d := Node{
		num: x,
		minNum:x,
	}
	if len(this.stack) > 0 && this.stack[len(this.stack)-1].minNum < x{
		d.minNum = this.stack[len(this.stack)-1].minNum
	}
	this.stack = append(this.stack, d)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].num
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].minNum
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */