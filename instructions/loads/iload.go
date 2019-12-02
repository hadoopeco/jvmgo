package loads

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
 * Date: 2018/3/27 18:15
 */

 // Load int from local variable

func _iload(frame *rtdata.Frame, index uint)  {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD_0 struct { base.NoOperandsInstruction }
func (self *ILOAD_0) Execute(frame *rtdata.Frame){
	_iload(frame, 0)
}

type ILOAD_1 struct { base.NoOperandsInstruction }

 func (self *ILOAD_1) Execute(frame *rtdata.Frame){
	_iload(frame, 1)
 }

 type ILOAD_2 struct { base.NoOperandsInstruction }
 func (self *ILOAD_2) Execute(frame *rtdata.Frame){
	_iload(frame, 2)
 }

 type ILOAD_3 struct { base.NoOperandsInstruction }
 func (self *ILOAD_3) Execute(frame *rtdata.Frame){
 	_iload(frame, 3)
 }

 type ILOAD_4 struct { base.NoOperandsInstruction }
 func (self *ILOAD_4) Execute(frame *rtdata.Frame){
 	_iload(frame, 4)
 }

type ILOAD struct { base.Index8Instruction }
 func (self *ILOAD) Execute(frame *rtdata.Frame){
 	_iload(frame, self.Index)
 }



