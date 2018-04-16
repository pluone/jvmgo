package main

import "fmt"
import "strings"
import "os"
import "jvmgo/classpath"
import "jvmgo/classfile"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1 ")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
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
	printClassInfo(cf)

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
