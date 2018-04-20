package rtda

type Thread struct {
	pc    int
	stack *JVMStack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) NewFrame(maxLocals, maxStack uint16) *Frame {
	return NewFrame(thread, maxLocals, maxStack)
}

func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}
