package rtdata

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:14
 */

type Frame struct {
	lower  *Frame  //链表数据结构
	localVars  LocalVars // 局部变量表指针
	operandStack  *OperandStack //操作数栈指针
	thread        *Thread       //线程指针
	nextPC        int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame{
	return &Frame{
		thread: thread,
		localVars:  newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) Thread() *Thread {
	return self.thread
}



