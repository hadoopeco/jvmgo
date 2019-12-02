package math

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
 * Date: 2018/3/27 19:41
 */

/**
* 布尔运算指令只能操作int 和 long变量 分为按位与(and) 按位或(or), 按位异或(xor)三种
*
 */
 type IAND struct { base.NoOperandsInstruction }

 func (self *IAND) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
 }

 type LAND struct {	base.NoOperandsInstruction }
 func (self *LAND) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
 }

