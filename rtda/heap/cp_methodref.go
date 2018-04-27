package heap

import (
	"jvmgo/classfile"
)

//MethodRef 方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, cfMethodRef *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.cp = cp
	methodRef.copyMemberRefInfo(&cfMethodRef.ConstantMemberInfo)
	return methodRef
}
