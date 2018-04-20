package rtda

//Frame JVMStack中的基本单元栈帧，栈帧中包括局部变量表和操作数栈
type Frame struct {
	lower        *Frame
	localVar     LocalVar
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint16) *Frame {
	return &Frame{
		thread:       thread,
		localVar:     newLocalVal(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (frame *Frame) LocalVar() LocalVar {
	return frame.localVar
}

func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(pc int) {
	frame.nextPC = pc
}

func (frame *Frame) Thread() *Thread {
	return frame.thread
}
