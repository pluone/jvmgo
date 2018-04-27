package base

import (
	"jvmgo/rtda"
)

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

//NoOperandsInstruction 表示无操作数的指令
type NoOperandsInstruction struct{}

func (nop *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	//nothing to do
}

//BranchInstruction 跳转指令
type BranchInstruction struct {
	Offset int
}

func (branch *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	val := reader.ReadInt16()
	branch.Offset = int(val)
}

//Index8Instruction 操作数为单字节的指令
type Index8Instruction struct {
	Index uint8
}

func (ins *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	ins.Index = reader.ReadUint8()
}

//Index16Instruction 操作数为两个字节的指令
type Index16Instruction struct {
	Index uint16
}

func (ins *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	ins.Index = reader.ReadUint16()
}
