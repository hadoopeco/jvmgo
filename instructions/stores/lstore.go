package stores

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
 * Date: 2018/3/27 18:29
 */


 // Store long into local variable

type LSTORE struct { base.Index8Instruction }
func (self *LSTORE) Execute(frame *rtdata.Frame){
	_lstore(frame, uint(self.Index))
}

type LSTORE_2 struct { base.Index8Instruction }
func (self *LSTORE_2) Execute(frame *rtdata.Frame){
	_lstore(frame, 2)
}


type LSTORE_0 struct { base.Index8Instruction }
func (self *LSTORE_0) Execute(frame *rtdata.Frame){
	_lstore(frame, 0)
}

type LSTORE_1 struct { base.Index8Instruction }
func (self *LSTORE_1) Execute(frame *rtdata.Frame){
	_lstore(frame, 1)
}



type LSTORE_3 struct { base.Index8Instruction }
func (self *LSTORE_3) Execute(frame *rtdata.Frame){
	_lstore(frame, 3)
}

func _lstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}