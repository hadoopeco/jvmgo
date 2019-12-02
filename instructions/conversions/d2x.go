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
 * Date: 2018/3/27 19:57

类型转换对应java中的基本类型强制转换操作
按照被转换变量的类型，类型转换指令可以分为3种：
i2x 系列指令把int变量强制转换成其它类型，
l2x系列指令把long变量强制转换成其它类型
f2x系列指令把float变量强制转换成其它类型
d2x系列之列把double变量强制转换成其它类型

*/

 type D2I struct{ base.NoOperandsInstruction }


 func (self *D2I) Execute(frame *rtdata.Frame){
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
 }

 type D2F struct{	base.NoOperandsInstruction }
 func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
 }

 type D2L struct{ base.NoOperandsInstruction }
 func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
 }



