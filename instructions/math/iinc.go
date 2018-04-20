package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (inc *IINC) FetchOperands(reader *base.ByteCodeReader) {
	inc.Index = uint(reader.ReadUint8())
	inc.Const = int32(reader.ReadInt8())
}

func (inc *IINC) Execute(frame *rtda.Frame) {
	localVals := frame.LocalVar()
	val := localVals.GetInt(uint(inc.Index))
	val += inc.Const
	localVals.SetInt(inc.Index, val)
}
