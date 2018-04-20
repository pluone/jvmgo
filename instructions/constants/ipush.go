package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (ins *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	ins.val = reader.ReadInt8()
}

func (ins *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(ins.val))
}

func (ins *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	ins.val = reader.ReadInt16()
}

func (ins *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(ins.val))
}
