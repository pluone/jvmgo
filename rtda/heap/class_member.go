package heap

import (
	"jvmgo/classfile"
)

//ClassMember 类成员包括属性和方法
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (classMember *ClassMember) Class() *Class {
	return classMember.class
}

func (classMember *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	classMember.accessFlags = memberInfo.AccessFlags()
	classMember.name = memberInfo.Name()
	classMember.descriptor = memberInfo.Descriptor()
}

func (classMember *ClassMember) isAccessibleTo(other *Class) bool {
	if classMember.IsPublic() {
		return true
	}

	thisClass := classMember.class
	if classMember.IsProtected() {
		return thisClass == other || other.IsSubClassOf(thisClass) || thisClass.getPackageName() == other.getPackageName()
	}

	if !classMember.IsPrivate() {
		return thisClass.getPackageName() == other.getPackageName()
	}
	return thisClass == other
}

func (classMember *ClassMember) IsStatic() bool {
	return 0 != classMember.accessFlags&ACC_STATIC
}

func (classMember *ClassMember) IsPublic() bool {
	return 0 != classMember.accessFlags&ACC_PUBLIC
}

func (classMember *ClassMember) IsProtected() bool {
	return 0 != classMember.accessFlags&ACC_PROTECTED
}

func (classMember *ClassMember) IsPrivate() bool {
	return 0 != classMember.accessFlags&ACC_PRIVATE
}

func (classMember *ClassMember) Name() string {
	return classMember.name
}

func (classMember *ClassMember) Descriptor() string {
	return classMember.descriptor
}
