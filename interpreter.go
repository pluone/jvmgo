package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/instructions"
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttribute := methodInfo.CodeAttribute()
	println("------------")
	println(codeAttribute)
	maxLocals := codeAttribute.MaxLocals()
	maxStack := codeAttribute.MaxStack()
	bytecode := codeAttribute.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {//todo printf中的格式化参数
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVar())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(pc, bytecode)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
