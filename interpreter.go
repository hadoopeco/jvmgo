package main

import (
	"fmt"
	"jvmgo/instructions"
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/28 2:31
 */

func interpret(method *heap.Method) {

	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	fmt.Printf("interpret :%v \n", method.Code())
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v \n", frame.LocalVars())
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
		fmt.Printf("pc:%2d opcode:0x%2d inst:%T %v\n", pc, opcode, inst, inst)
		inst.Execute(frame)
	}
}
