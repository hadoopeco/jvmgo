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
 * Date: 2018/3/28 10:21
 */
/*
   store 存储类型指令类似于load加载指令
*/
// Store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (self *ASTORE) Execute(frame *rtdata.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *rtdata.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *rtdata.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *rtdata.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *rtdata.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtdata.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}