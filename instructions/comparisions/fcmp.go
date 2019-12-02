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
 * Date: 2018/3/27 20:10
 */
 /**
 还有第四种结果无法比较，fcmpg 和fcmpgl指令区别就在于对第四种结果的定义
 */

type FCMPG struct { base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtdata.Frame){
	_fcmp(frame, true)
}

type FCMPL struct {	 base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtdata.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtdata.Frame, gFlag bool){
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else  if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}

}


