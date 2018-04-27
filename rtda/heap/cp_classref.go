package heap

import (
	"jvmgo/classfile"
)

//ClassRef 类符号引用
type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, cfClassInfo *classfile.ConstantClassInfo) *ClassRef {
	classRef := &ClassRef{}
	classRef.cp = cp
	classRef.className = cfClassInfo.Value()
	return classRef
}
