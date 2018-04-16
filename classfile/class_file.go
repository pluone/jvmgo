package classfile

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool //todo 都换成指针类型可以吗
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

//Parse parse ClassData to ClassFile
func Parse(classData []byte) (*ClassFile, error) {
	//todo 原代码中有recover()后续考虑
	cr := &ClassReader{classData}
	cf := ClassFile{}
	cf.read(cr)
	return &cf, nil //todo
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, &cf.constantPool)
	cf.methods = readMembers(reader, &cf.constantPool)
	cf.attributes = readAttributes(reader, &cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(cr *ClassReader) {
	cf.magic = cr.readUint32()
	if cf.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(cr *ClassReader) {
	cf.minorVersion = cr.readUint16()
	cf.majorVersion = cr.readUint16()

	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) ThisClass() uint16 {
	return cf.thisClass
}

func (cf *ClassFile) SuperClass() uint16 {
	return cf.superClass
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, classIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(classIndex)
	}
	return interfaceNames
}

func (cf *ClassFile) Fields() []*MemberInfo{
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo{
	return cf.methods
}
