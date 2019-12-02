package main

import (
	"jvmgo/classfile"
	"jvmgo/rtdata"
	"jvmgo/instructions/base"
	"jvmgo/instructions"
	"fmt"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/28 2:31
 */

func interpret(methodInfo *classfile.MemberInfo)  {
	codeAttr := methodInfo.CodeAttribute() //获取code属性
	maxLoacls := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack() //获取操作数空间
	bytecode := codeAttr.Code()  //获取字节码
	for c := range bytecode{
		fmt.Printf("0x%x ",bytecode[c])
	}
	fmt.Println(" ")
	thread := rtdata.NewThread()
	frame := thread.NewFrame(maxLoacls,  maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}


func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v \n", frame.LocalVars() )
		fmt.Printf("OperandsStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

//循环执行 "计算pc, 解码指令, 执行指令"
func loop(thread *rtdata.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d nextpc:%2d opcode:0x%x inst:%T %v\n", pc, reader.PC(), opcode, inst, inst )
		inst.Execute(frame)
	}
}



