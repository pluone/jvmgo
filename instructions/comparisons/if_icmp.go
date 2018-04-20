package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }

func (cmp *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 == val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func (cmp *IF_ICMPNE) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 != val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func (cmp *IF_ICMPGT) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 > val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func (cmp *IF_ICMPGE) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func (cmp *IF_ICMPLT) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 < val2 {
		base.Branch(frame, cmp.Offset)
	}
}

func (cmp *IF_ICMPLE) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, cmp.Offset)
	}
}
