package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type LSTORE struct{ base.Index8Instruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVar().SetLong(index, val)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVar().SetInt(index, val)
}

func (store *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (store *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (store *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(store.Index))
}

func (store *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (store *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (store *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (store *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
