package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (dup *DUP) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot)
	frame.OperandStack().PushSlot(slot)
}


