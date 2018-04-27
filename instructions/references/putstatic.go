package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type PutStatic struct{ base.Index16Instruction }

func (putStatic *PutStatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	constantPool := currentClass.ConstantPool()
	fieldRef := constantPool.GetConstant(uint(putStatic.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !class.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if class.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError!")
		}
	}
	descriptor := field.Descriptor()
	staticVars := class.StaticVars()
	slotId := field.SlotId()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		staticVars.SetInt(slotId, frame.OperandStack().PopInt())
	case 'F':
		staticVars.SetFloat(slotId, frame.OperandStack().PopFloat())
	case 'D':
		staticVars.SetDouble(slotId, frame.OperandStack().PopDouble())
	case 'J':
		staticVars.SetLong(slotId, frame.OperandStack().PopLong())
	case 'L', '[': //引用类型和数组类型
		staticVars.SetRef(slotId, frame.OperandStack().PopRef())
	default:
		panic("unknown field descriptor")
	}
}
