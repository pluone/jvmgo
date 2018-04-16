package classfile

//MemberInfo MethodInfo和Field的抽象
type MemberInfo struct {
	constantPool    *ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(cr *ClassReader, cp *ConstantPool) []*MemberInfo {
	count := cr.readUint16()
	members := make([]*MemberInfo, count)
	for i := range members {
		members[i] = readMember(cr, cp)
	}
	return members
}

func readMember(cr *ClassReader, cp *ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (member *MemberInfo) Name() string {
	return member.constantPool.getUtf8String(member.nameIndex)
}
