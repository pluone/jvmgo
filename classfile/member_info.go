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

func (member *MemberInfo) AccessFlags() uint16 {
	return member.accessFlags
}

func (member *MemberInfo) Name() string {
	return member.constantPool.getUtf8String(member.nameIndex)
}

func (member *MemberInfo) Descriptor() string {
	return member.constantPool.getUtf8String(member.descriptorIndex)
}

func (member *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range member.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (member *MemberInfo) ConstatnValueAttribute() *ConstantValueAttribute {
	for _, attr := range member.attributes {
		switch attr.(type) {
		case *ConstantValueAttribute:
			return attr.(*ConstantValueAttribute)
		}
	}
	return nil
}
