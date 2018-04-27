package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type GetField struct {
	base.Index16Instruction
}

func (getField *GetField) Execute(frame rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	fieldRef := constantPool.GetConstant(uint(getField.Index)).(heap.FieldRef)
	field := fieldRef.ResolvedField()
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	operandStack := frame.OperandStack()
	ref := operandStack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	instanceFields := ref.InstanceFields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		operandStack.PushInt(instanceFields.GetInt(slotId))
	case 'F':
		operandStack.PushFloat(instanceFields.GetFloat(slotId))
	case 'D':
		operandStack.PushDouble(instanceFields.GetDouble(slotId))
	case 'J':
		operandStack.PushLong(instanceFields.GetLong(slotId))
	case 'L', '[': //引用类型和数组类型
		operandStack.PushRef(instanceFields.GetRef(slotId))
	}

}
