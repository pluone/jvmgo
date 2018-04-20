package base

//ByteCodeReader 字节码读取器
type ByteCodeReader struct {
	code []byte
	pc   int
}

func (reader *ByteCodeReader) ReadUint8() uint8 {
	val := reader.code[reader.pc]
	reader.pc++
	return val
}

func (reader *ByteCodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *ByteCodeReader) ReadUint16() uint16 {
	high := reader.ReadUint8()
	low := reader.ReadUint8()
	return uint16(low) | uint16(high)<<8
}

func (reader *ByteCodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *ByteCodeReader) ReadInt32() int32 {
	byte1 := reader.ReadUint8()
	byte2 := reader.ReadUint8()
	byte3 := reader.ReadUint8()
	byte4 := reader.ReadUint8()
	return int32(uint32(byte4)<<24 | uint32(byte3)<<16 | uint32(byte2)<<8 | uint32(byte1))
}

func (reader *ByteCodeReader) PC() int {
	return reader.pc
}

func (reader *ByteCodeReader) Reset(pc int, bytecode []byte) {
	reader.pc = pc
	reader.code = bytecode
}
