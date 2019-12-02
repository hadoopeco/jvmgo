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
 * Date: 2018/3/28 21:05
 */

 type DCMPG struct { base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtdata.Frame)  {
	_dcmp(frame, true)
}

type DCMPL struct {	base.NoOperandsInstruction }
func (self *DCMPL) Execute(frame *rtdata.Frame)  {
	_dcmp(frame, false)
}

func _dcmp(frame *rtdata.Frame, gFlag bool){
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(2)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}

}