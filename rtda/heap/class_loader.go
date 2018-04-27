package heap

import (
	"jvmgo/classfile"
	"jvmgo/classpath"
)

//ClassLoader ClassLoader
type ClassLoader struct {
	classpath *classpath.Classpath
	classMap  map[string]*Class //loaded class
}

//NewClassLoader 创新一个新的ClassLoader
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		classpath: cp,
		classMap:  make(map[string]*Class),
	}
}

func (cl *ClassLoader) LoadClass(className string) *Class {
	if class, ok := cl.classMap[className]; ok {
		return class
	}
	return cl.loadNonArrayClass(className)
}

func (cl *ClassLoader) loadNonArrayClass(className string) *Class {
	data, _ := cl.readClass(className)
	class := cl.defineClass(data)
	link(class)
	return class
}

func (cl *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := cl.classpath.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException " + className)
	}
	return data, entry
}

func (cl *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = cl
	resolveSuperClass(class)
	resolveInterfaces(class)
	cl.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceNames := class.interfaceNames
	class.interfaces = make([]*Class, len(interfaceNames))
	for i, interfaceName := range interfaceNames {
		class.interfaces[i] = class.loader.LoadClass(interfaceName)
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//empty 验证的步骤忽略
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	slotId := field.slotId
	descriptor := field.descriptor
	if constVal := class.constantPool.GetConstant(field.constValueIndex); constVal != nil {
		switch descriptor {
		case "Z", "B", "C", "S", "I":
			class.staticVars.SetInt(slotId, constVal.(int32))
		case "J":
			class.staticVars.SetLong(slotId, constVal.(int64))
		case "F":
			class.staticVars.SetFloat(slotId, constVal.(float32))
		case "D":
			class.staticVars.SetDouble(slotId, constVal.(float64))
		case "Ljava/lang/String":
			panic("todo")
		}
	}
}
