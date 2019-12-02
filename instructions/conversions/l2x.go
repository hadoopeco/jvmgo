package conversions

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/28 23:51
 */

// Convert long to double
type L2D struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct{ base.NoOperandsInstruction }

func (self *L2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct{ base.NoOperandsInstruction }

func (self *L2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}