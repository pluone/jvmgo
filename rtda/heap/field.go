package heap

import (
	"jvmgo/classfile"
)

//Field 代表类的属性
type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	classFields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		classFields[i] = &Field{}
		classFields[i].class = class
		classFields[i].copyMemberInfo(cfField)
		classFields[i].copyAttribute(cfField)
	}
	return classFields
}

func (field *Field) IsStatic() bool {
	return 0 != field.accessFlags&ACC_STATIC
}

func (field *Field) IsFinal() bool {
	return 0 != field.accessFlags&ACC_FINAL
}

func (field *Field) IsLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}

func (field *Field) copyAttribute(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstatnValueAttribute(); valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field *Field) SlotId() uint{
	return  field.slotId
}
