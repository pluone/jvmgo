package heap

import (
	"jvmgo/classfile"
)

//InterfaceMethodRef 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newIntefaceMethodRef(cp *ConstantPool, cfIntefaceMethodRef *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	imref := &InterfaceMethodRef{}
	imref.cp = cp
	imref.copyMemberRefInfo(&cfIntefaceMethodRef.ConstantMemberInfo)
	return imref
}
