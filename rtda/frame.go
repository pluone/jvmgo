package rtda

import "jvmgo/rtda/heap"

//Frame JVMStack中的基本单元栈帧，栈帧中包括局部变量表和操作数栈
type Frame struct {
	lower        *Frame
	localVar     LocalVar
	operandStack *OperandStack
	method       *heap.Method
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		method:       method,
		thread:       thread,
		localVar:     newLocalVal(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
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

func (frame *Frame) Method() *heap.Method {
	return frame.method
}
