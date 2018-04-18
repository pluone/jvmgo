package rtda

//Frame JVMStack中的基本单元栈帧，栈帧中包括局部变量表和操作数栈
type Frame struct {
	lower        *Frame
	localVar    LocalVar
	operandStack *OperandStack
}

//NewFrame 创建栈帧
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVar:     newLocalVal(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (frame *Frame) LocalVar() LocalVar {
	return frame.localVar
}

func (frame *Frame) OperandStack() *OperandStack{
	return frame.operandStack
}
