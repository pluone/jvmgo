package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type ALOAD struct{ base.Index8Instruction }
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func _ALOAD(frame *rtda.Frame, index uint) {
	val := frame.LocalVar().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (load *ALOAD) Execute(frame *rtda.Frame) {
	_ALOAD(frame, uint(load.Index))
}

func (load *ALOAD_0) Execute(frame *rtda.Frame) {
	_ALOAD(frame, 0)
}

func (load *ALOAD_1) Execute(frame *rtda.Frame) {
	_ALOAD(frame, 1)
}

func (load *ALOAD_2) Execute(frame *rtda.Frame) {
	_ALOAD(frame, 2)
}

func (load *ALOAD_3) Execute(frame *rtda.Frame) {
	_ALOAD(frame, 3)
}
