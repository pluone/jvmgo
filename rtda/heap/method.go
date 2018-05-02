package heap

import (
	"jvmgo/classfile"
)

//Method 代表Class中的方法
type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
}

func (method *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute();codeAttr!=nil{
		method.maxLocals = uint(codeAttr.MaxLocals())
		method.maxStack = uint(codeAttr.MaxStack())
		method.code = codeAttr.Code()
	}
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
	}
	return methods
}

func (method *Method) MaxLocals() uint {
	return method.maxLocals
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}

func (method *Method) Code() []byte {
	return method.code
}
