package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type PutField struct {
	base.Index16Instruction
}

func (putField *PutField) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	constantPool := currentClass.ConstantPool()
	fieldRef := constantPool.GetConstant(uint(putField.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError!")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	operandStack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := operandStack.PopInt()
		object := operandStack.PopRef()
		if object == nil {
			panic("java.lang.NullPointerException!")
		}
		object.InstanceFields().SetInt(slotId, val)
	case 'F':
		val := operandStack.PopFloat()
		object := operandStack.PopRef()
		if object == nil {
			panic("java.lang.NullPointerException!")
		}
		object.InstanceFields().SetFloat(slotId, val)
	case 'D':
		val := operandStack.PopDouble()
		object := operandStack.PopRef()
		if object == nil {
			panic("java.lang.NullPointerException!")
		}
		object.InstanceFields().SetDouble(slotId, val)
	case 'J':
		val := operandStack.PopLong()
		object := operandStack.PopRef()
		if object == nil {
			panic("java.lang.NullPointerException!")
		}
		object.InstanceFields().SetLong(slotId, val)
	case 'L', '[': //引用类型和数组类型
		val := operandStack.PopRef()
		object := operandStack.PopRef()
		if object == nil {
			panic("java.lang.NullPointerException!")
		}
		object.InstanceFields().SetRef(slotId, val)
	}
}
