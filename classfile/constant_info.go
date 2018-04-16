package classfile

import (
	"fmt"
	"math"
	"strconv"
)

const (
	constUtf8               = 1
	constInteger            = 3
	constFloat              = 4
	constLong               = 5
	constDouble             = 6
	constClass              = 7
	constString             = 8
	constFieldRef           = 9
	constMehtodRef          = 10
	constInterfaceMethodRef = 11
	constNameAndType        = 12
	constMethodHandle       = 15
	constMethodType         = 16
	constInvokeDynamic      = 18
)

//ConstantInfo 接口
type ConstantInfo interface {
	readInfo(cr *ClassReader)
	String() string
}

func readConstantInfo(cr *ClassReader) ConstantInfo {
	tag := cr.readUint8()
	ci := newConstantInfo(tag)
	ci.readInfo(cr)
	return ci
}

func newConstantInfo(tag uint8) ConstantInfo {
	switch tag {
	case constUtf8:
		return &ConstantUtf8Info{}
	case constInteger:
		return &ConstantIntegerInfo{}
	case constFloat:
		return &ConstantFloatInfo{}
	case constLong:
		return &ConstantLongInfo{}
	case constDouble:
		return &ConstantDoubleInfo{}
	case constClass:
		return &ConstantClassInfo{}
	case constString:
		return &ConstantStringInfo{}
	case constNameAndType:
		return &ConstantNameAndTypeInfo{}
	case constFieldRef:
		return &ConstantFieldRefInfo{}
	case constMehtodRef:
		return &ConstantMethodRefInfo{}
	case constInterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{}
	default:
		fmt.Printf("constant tag is: %v", tag)
		panic("invalid constant type!")
	}
}

//ConstantUtf8Info utf8类型常量
type ConstantUtf8Info struct {
	str string
}

func (utf8Info *ConstantUtf8Info) readInfo(cr *ClassReader) {
	length := uint32(cr.readUint16())
	bytes := cr.readBytes(length)
	utf8Info.str = decodeMUTF8(bytes)
}

func (utf8Info *ConstantUtf8Info) String() string {
	return utf8Info.str
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes) //todo 简化版本
}

//ConstantIntegerInfo integer类型常量
type ConstantIntegerInfo struct {
	val int32
}

func (intInfo *ConstantIntegerInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	intInfo.val = int32(bytes)
}

func (intInfo *ConstantIntegerInfo) String() string {
	return strconv.Itoa(int(intInfo.val))
}

//ConstantFloatInfo 浮点数常量
type ConstantFloatInfo struct {
	val float32
}

func (floatInfo *ConstantFloatInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint32()
	floatInfo.val = math.Float32frombits(bytes)
}

func (floatInfo *ConstantFloatInfo) String() string {
	return strconv.FormatFloat(float64(floatInfo.val), 'e', -1, 64)
}

//ConstantLongInfo 长整型常量
type ConstantLongInfo struct {
	val int64
}

func (longInfo *ConstantLongInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	longInfo.val = int64(bytes)
}

func (longInfo *ConstantLongInfo) String() string {
	return strconv.FormatInt(longInfo.val, 10)
}

//ConstantDoubleInfo 双精度浮点型常量
type ConstantDoubleInfo struct {
	val float64
}

func (doubleInfo *ConstantDoubleInfo) readInfo(cr *ClassReader) {
	bytes := cr.readUint64()
	doubleInfo.val = math.Float64frombits(bytes)
}

func (doubleInfo *ConstantDoubleInfo) String() string {
	return strconv.FormatFloat(doubleInfo.val, 'e', -1, 64)
}

//ConstantStringInfo 字符串类型常量
type ConstantStringInfo struct {
	stringIndex uint16
}

func (stringInfo *ConstantStringInfo) readInfo(cr *ClassReader) {
	stringInfo.stringIndex = cr.readUint16()
}

func (stringInfo *ConstantStringInfo) String() string {
	return "constant string #" + strconv.Itoa(int(stringInfo.stringIndex))
}

//ConstantClassInfo class常量
type ConstantClassInfo struct {
	classNameIndex uint16
}

func (classInfo *ConstantClassInfo) readInfo(cr *ClassReader) {
	classInfo.classNameIndex = cr.readUint16()
}

func (classInfo *ConstantClassInfo) String() string {
	return "constant class #" + strconv.Itoa(int(classInfo.classNameIndex))
}

//ConstantNameAndTypeInfo nameAndType类型常量
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (info *ConstantNameAndTypeInfo) readInfo(cr *ClassReader) {
	info.nameIndex = cr.readUint16()
	info.descriptorIndex = cr.readUint16()
}

func (info *ConstantNameAndTypeInfo) String() string {
	return "constant name and index, nameIndex: #" + strconv.Itoa(int(info.nameIndex)) + " descriptorIndex: #" + strconv.Itoa(int(info.descriptorIndex))
}

type ConstantMemberInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberInfo) readInfo(cr *ClassReader) {
	info.classIndex = cr.readUint16()
	info.nameAndTypeIndex = cr.readUint16()
}

func (info *ConstantMemberInfo) String() string {
	return "constant member, classIndex: #" + strconv.Itoa(int(info.classIndex)) + " nameAndTypeIndex: #" + strconv.Itoa(int(info.nameAndTypeIndex))
}

type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}
