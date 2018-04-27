package heap

import (
	"jvmgo/classfile"
	"fmt"
)

//Method 代表Class中的方法
type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
}

func (method *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	fmt.Printf("cfMethod is %#v\n",cfMethod)
	method.maxLocals = uint(cfMethod.CodeAttribute().MaxLocals())
	method.maxStack = uint(cfMethod.CodeAttribute().MaxStack())
	method.code = cfMethod.CodeAttribute().Code()
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	fmt.Printf("cfMethods: %#v\n",cfMethods)
	fmt.Printf("cfMethods count: %#v\n",len(cfMethods))
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		fmt.Printf("class file method,method name:%s\n %#v\n",cfMethod.Name(),cfMethod)
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
