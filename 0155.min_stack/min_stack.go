package min_stack

type MinStack struct {
	stack [][2]int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([][2]int, 0),
	}
}

func (ms *MinStack) Push(x int) {
	var min int
	if ms.Empty() {
		min = x
	} else {
		min = ms.GetMin()
		if x < min {
			min = x
		}
	}
	ms.stack = append(ms.stack, [2]int{x, min})
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[0 : len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1][0]
}

func (ms *MinStack) Empty() bool {
	return len(ms.stack) == 0
}

func (ms *MinStack) GetMin() int {
	return ms.stack[len(ms.stack)-1][1]
}
