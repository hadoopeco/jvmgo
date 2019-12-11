package rtdata

import "jvmgo/rtdata/heap"

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:07


	JVM
	  Thread
		pc
		Stack
		  Frame
			LocalVars
			OperandStack
*/

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	//Frame.NewFrame
	return newFrame(self, method)
}

func (self *Thread) Stack() *Stack {
	return self.stack
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
