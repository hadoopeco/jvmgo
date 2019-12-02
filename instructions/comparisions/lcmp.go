package comparisions

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
 * Date: 2018/3/27 20:07
 */

type LCMP struct {	base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtdata.Frame){
 	stack := frame.OperandStack()
 	v2 := stack.PopLong()
 	v1 := stack.PopLong()

 	if v1 > v2 {
 		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
 }