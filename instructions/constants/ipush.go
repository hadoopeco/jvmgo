package constants

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
 * Date: 2018/3/27 18:11
 */

// Push byte
 type BIPUSH struct { val int8 }

 func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt8()
 }

 func (self *BIPUSH) Execute(frame *rtdata.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
 }

// Push short
 type SIPUSH struct { val int16 }
 func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt16()
 }

 func (self *SIPUSH) Execute(frame *rtdata.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
 }

