package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type DADD struct{ base.NoOperandsInstruction }

func (add *IADD) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopInt()
	v1 := operandStack.PopInt()
	result := v1 + v2
	operandStack.PushInt(result)
}

func (add *LADD) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopLong()
	v1 := operandStack.PopLong()
	result := v1 + v2
	operandStack.PushLong(result)
}

func (add *FADD) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopFloat()
	v1 := operandStack.PopFloat()
	result := v1 + v2
	operandStack.PushFloat(result)
}

func (add *DADD) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopDouble()
	v1 := operandStack.PopDouble()
	result := v1 + v2
	operandStack.PushDouble(result)
}
