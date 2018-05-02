package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type InvokeSpecial struct {
	base.Index16Instruction
}

func (invokeSpecial *InvokeSpecial) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PopRef()//todo
}