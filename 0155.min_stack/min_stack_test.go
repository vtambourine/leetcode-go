package min_stack

import "testing"

type step struct {
	method string
	value  int
}

type test struct {
	steps []step
}

func TestMinStack(t *testing.T) {
	tests := []test{
		{[]step{
			{"push", -2},
			{"push", 0},
			{"push", -3},
			{"getMin", -3},
			{"pop", 0},
			{"top", 0},
			{"getMin", -2},
			{"pop", 0},
			{"getMin", -2},
		},
		},
	}

	for _, c := range tests {
		ms := Constructor()
		for _, s := range c.steps {
			switch s.method {
			case "push":
				ms.Push(s.value)
			case "pop":
				ms.Pop()
			case "top":
				top := ms.Top()
				if top != s.value {
					t.Fatalf("testing %v\n(*MinStack).Top() fails.\nExpected %v\nReceived %v", c.steps, s.value, top)
				}
			case "getMin":
				min := ms.GetMin()
				if min != s.value {
					t.Fatalf("testing %v\n(*MinStack).GetMin() fails.\nExpected %v\nReceived %v", c.steps, s.value, min)
				}
			}
		}
	}
}
