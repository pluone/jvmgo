package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type LDC struct {
	base.Index8Instruction
}

type LDC_W struct {
	base.Index16Instruction
}

type LDC2_W struct {
	base.Index16Instruction
}

func (ldc *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, uint(ldc.Index))
}

func (ldc *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, uint(ldc.Index))
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	constantPool := frame.Method().Class().ConstantPool()
	constant := constantPool.GetConstant(index)
	switch constant.(type) {
	case int32:
		stack.PushInt(constant.(int32))
	case float32:
		stack.PushFloat(constant.(float32))
	default:
		panic("todo ldc")
	}
}

func (ldc *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	constantPool := frame.Method().Class().ConstantPool()
	constant := constantPool.GetConstant(uint(ldc.Index))
	switch constant.(type) {
	case int64:
		stack.PushLong(constant.(int64))
	case float64:
		stack.PushDouble(constant.(float64))
	default:
		panic("todo ldc")
	}
}
