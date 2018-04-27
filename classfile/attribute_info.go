package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(cr *ClassReader, cp *ConstantPool) []AttributeInfo {
	count := cr.readUint16()
	attributes := make([]AttributeInfo, count)
	for i := range attributes {
		attributes[i] = readAttribute(cr, cp)
	}
	return attributes
}

func readAttribute(cr *ClassReader, cp *ConstantPool) AttributeInfo {
	attributeNameIndex := cr.readUint16()
	attrName := cp.getUtf8String(attributeNameIndex)
	attributeLength := cr.readUint32()
	attrInfo := newAttributeInfo(attrName, attributeLength, cp)
	attrInfo.readInfo(cr)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLength uint32, cp *ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{constantPool: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	default:
		return &UnparsedAttributeInfo{attrName, attrLength, nil}
	}
}

type UnparsedAttributeInfo struct {
	attrName   string
	attrLength uint32
	attrInfo   []byte
}

func (attr *UnparsedAttributeInfo) readInfo(reader *ClassReader) {
	attr.attrInfo = reader.readBytes(attr.attrLength)
}

//CodeAttribute methodInfo中的CodeAttribute
type CodeAttribute struct {
	constantPool   *ConstantPool
	maxLocals      uint16
	maxStack       uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

//ExceptionTableEntry codeAttribute中的异常表项
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	attr.code = reader.readBytes(codeLength)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.constantPool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (attr *CodeAttribute) MaxLocals() uint16 {
	return attr.maxLocals
}

func (attr *CodeAttribute) MaxStack() uint16 {
	return attr.maxStack
}

func (attr *CodeAttribute) Code() []byte {
	return attr.code
}
