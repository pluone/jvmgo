package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"math"
)

type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type DREM struct{ base.NoOperandsInstruction }

func (rem *IREM) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	v1 := operandStack.PopInt()
	result := v1 % v2
	operandStack.PushInt(result)
}

func (rem *LREM) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero!")
	}
	v1 := operandStack.PopLong()
	result := v1 % v2
	operandStack.PushLong(result)
}

func (rem *FREM) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopFloat()
	v1 := operandStack.PopFloat()
	result := math.Mod(float64(v1), float64(v2))
	operandStack.PushFloat(float32(result))
}

func (rem *DREM) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	v2 := operandStack.PopDouble()
	v1 := operandStack.PopDouble()
	result := math.Mod(v1, v2)
	operandStack.PushDouble(result)
}
