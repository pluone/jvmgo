package rtda

type JVMStack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *JVMStack {
	return &JVMStack{
		maxSize: maxSize,
	}
}

func (stack *JVMStack) push(frame *Frame) {
	if stack.size+1 > stack.maxSize {
		panic("java.lang.StackOverflowException!")
	}

	frame.lower = stack._top
	stack._top = frame
	stack.size++
}

func (stack *JVMStack) pop() *Frame {
	if stack.size == 0 {
		panic("jvm stack is empty!")
	}
	frame := stack._top
	stack._top = frame.lower
	stack.size--
	return frame
}

func (stack *JVMStack) top() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}
	return stack._top
}
