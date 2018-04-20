package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (pop *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (pop *POP2) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	operandStack.PopSlot()
	operandStack.PopSlot()
}
