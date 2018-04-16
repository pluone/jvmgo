package classfile

type AttributeInfo struct {
	attributeNameIndex uint16
	attributeLength    uint32
	attributeContent   []byte
}

func readAttributes(cr *ClassReader, cp *ConstantPool) []AttributeInfo {
	count := cr.readUint16()
	attributs := make([]AttributeInfo, count)
	for i := range attributs {
		attributs[i] = readAttribute(cr, cp)
	}
	return attributs
}

func readAttribute(cr *ClassReader, cp *ConstantPool) AttributeInfo {
	attributeNameIndex := cr.readUint16()
	attributeLength := cr.readUint32()
	attributeContent := cr.readBytes(attributeLength)
	return AttributeInfo{
		attributeNameIndex,
		attributeLength,
		attributeContent,
	}
}
