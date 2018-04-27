package heap

import (
	"jvmgo/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string //this class name
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlags&ACC_PUBLIC
}

func (class *Class) IsAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i > 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) IsSubClassOf(other *Class) bool {
	return class.superClass == other
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}

func (class *Class) IsInterface() bool {
	return 0 != class.accessFlags&ACC_INTERFACE
}

func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlags&ACC_ABSTRACT
}

func (class *Class) NewObject() *Object {
	return &Object{
		class:          class,
		instanceFields: newSlots(class.instanceSlotCount),
	}
}

func (class *Class) StaticVars() Slots {
	return class.staticVars
}

func (class *Class) IsStatic() bool {
	return 0 != class.accessFlags&ACC_STATIC
}

func (class *Class) IsFinal() bool {
	return 0 != class.accessFlags&ACC_FINAL
}
func (class *Class) GetMainMethod() *Method {
	for _,method := range class.methods{
		if method.Name() == "main" && method.Descriptor()=="([Ljava/lang/String;)V" && method.IsStatic(){
			return method
		}
	}
	return nil
}
