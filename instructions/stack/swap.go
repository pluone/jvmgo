package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (swap *SWAP) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot2)
}
