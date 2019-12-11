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

func interpret(method *heap.Method, logInst bool) {

	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	fmt.Printf("befor loop thread.stack = %v \n", thread.Stack())
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

//循环执行 "计算pc, 解码指令, 执行指令"
func loop(thread *rtdata.Thread, logInst bool) {

	reader := &base.BytecodeReader{}
	for {
		//frame := thread.PopFrame()
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		fmt.Printf("thread=%v code= %v current pc = %v \n", thread.Stack(), frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtdata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("logInstruction =%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
