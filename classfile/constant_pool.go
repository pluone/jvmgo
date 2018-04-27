package classfile

type ConstantPool []ConstantInfo

func readConstantPool(cr *ClassReader) ConstantPool {
	cpCount := int(cr.readUint16())
	cp := make([]ConstantInfo, cpCount)
	//The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(cr, cp)
		switch cp[i].(type) {
		case *ConstantDoubleInfo, *ConstantLongInfo:
			i++
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if val := cp[index]; val != nil {
		return val
	}
	panic("invalid constant pool index!")
}

func (cp ConstantPool) getClassName(nameIndex uint16) string {
	classInfo := cp.getConstantInfo(nameIndex).(*ConstantClassInfo)
	return cp.getUtf8String(classInfo.classNameIndex)
}

func (cp ConstantPool) getNameAndType(nameAndTypeIndex uint16) (string, string) {
	nameAndTypeInfo := cp.getConstantInfo(nameAndTypeIndex).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8String(nameAndTypeInfo.nameIndex)
	_type := cp.getUtf8String(nameAndTypeInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getUtf8String(index uint16) string {
	info := cp[index].(*ConstantUtf8Info)
	return info.str
}
