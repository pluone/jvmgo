package heap

import (
	"jvmgo/classfile"
)

//MemberRef 类成员（字段和方法）符号引用
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (memberRef *MemberRef) copyMemberRefInfo(memberInfo *classfile.ConstantMemberInfo) {
	memberRef.className = memberInfo.ClassName()
	memberRef.name, memberRef.descriptor = memberInfo.NameAndDescriptor()
}

func (memberRef *MemberRef) Name() string {
	return memberRef.name
}

func (memberRef *MemberRef) Descriptor() string {
	return memberRef.descriptor
}
