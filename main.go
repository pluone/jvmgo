package main

import "fmt"
import "strings"
import "os"
import "jvmgo/classpath"
import "jvmgo/classfile"
import "jvmgo/rtda"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1 ")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		testJVMStack()
		startJVM(cmd)
	}
}

func testJVMStack() {
	frame := rtda.NewFrame(nil, 100, 100)
	testLocalVals(frame.LocalVar())
	testOperandStack(frame.OperandStack())
}

func testLocalVals(vars rtda.LocalVar) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())

}

func startJVM(cmd *Cmd) {

	cpath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf(" classpath:%v class:%v args:%v\n", cpath, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", string(os.PathSeparator), -1)
	classData, _, err := cpath.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		fmt.Println("parse class data failed")
	}
	// printClassInfo(cf)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	}

}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	methods := cf.Methods()
	for _, method := range methods {
		if "main" == method.Name() && "([Ljava/lang/String;)V" == method.Descriptor() {
			return method
		}
	}
	return nil
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	constantPool := cf.ConstantPool()
	fmt.Printf("constants count: %v\n", len(constantPool))
	for i := 1; i < len(constantPool); i++ {
		fmt.Printf("\t#%v, value: %s\n", i, constantPool[i])
		switch constantPool[i].(type) {
		case *classfile.ConstantDoubleInfo, *classfile.ConstantLongInfo:
			i++
		}
	}
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ThisClass())
	fmt.Printf("super class: %v\n", cf.SuperClass())
	fmt.Printf("interfaces: %s\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, member := range cf.Fields() {
		fmt.Printf("\tfields name: %s\n", member.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, member := range cf.Methods() {
		fmt.Printf("\tmethods name: %s\n", member.Name())
	}
}
