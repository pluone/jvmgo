package heap

import (
	"jvmgo/classfile"
)

//Constant 运行时常量池中的常量
type Constant interface{}

//ConstantPool 运行时常量池
type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {
	consts := make([]Constant, len(cfcp))
	rtcp := &ConstantPool{class, consts}
	for i := 1; i < len(cfcp); i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.Value()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtcp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtcp, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtcp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			iMethodInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newIntefaceMethodRef(rtcp, iMethodInfo)
		}
	}
	return rtcp
}

//GetConstant 根据索引从常量池中获取一个常量
func (cp *ConstantPool) GetConstant(index uint) Constant {
	if constant := cp.consts[index]; constant != nil {
		return constant
	}
	return nil
}
