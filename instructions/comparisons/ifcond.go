package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}

type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}

func (cond *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, cond.Offset)
	}
}

func (cond *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, cond.Offset)
	}
}

func (cond *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, cond.Offset)
	}
}

func (cond *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, cond.Offset)
	}
}

func (cond *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, cond.Offset)
	}
}

func (cond *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, cond.Offset)
	}
}
