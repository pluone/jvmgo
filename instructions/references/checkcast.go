package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type CheckCast struct {
	base.Index16Instruction
}

func (checkCast *CheckCast) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	classRef := constantPool.GetConstant(uint(checkCast.Index)).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	ref := frame.OperandStack().PopRef()
	if ref.IsInstanceOf(class) {
		frame.OperandStack().PushInt(1)
	}else {
		frame.OperandStack().PushInt(0)
	}
}