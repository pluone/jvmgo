package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type GetStatic struct {
	base.Index16Instruction
}

func (getStatic *GetStatic) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	fieldRef := constantPool.GetConstant(uint(getStatic.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	slotId := field.SlotId()
	descriptor := field.Descriptor()
	staticVars := field.Class().StaticVars()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		frame.OperandStack().PushInt(staticVars.GetInt(slotId))
	case 'F':
		frame.OperandStack().PushFloat(staticVars.GetFloat(slotId))
	case 'D':
		frame.OperandStack().PushDouble(staticVars.GetDouble(slotId))
	case 'J':
		frame.OperandStack().PushLong(staticVars.GetLong(slotId))
	case 'L', '[': //引用类型和数组类型
		frame.OperandStack().PushRef(staticVars.GetRef(slotId))
	}
}
