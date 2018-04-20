package constants

import "jvmgo/instructions/base"
import "jvmgo/rtda"

//NOP nop instruction
type NOP struct {
	base.NoOperandsInstruction
}

//Execute nop
func (nop *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}

