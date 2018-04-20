package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVar().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (load *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(load.Index))
}

func (load *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (load *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (load *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (load *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
