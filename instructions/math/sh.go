package math

import (
	"jvmgo/rtdata"
	"jvmgo/instructions/base"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/27 19:23

*/


// int 左位移
type ISHL struct {	 base.NoOperandsInstruction }

 /* 先从操作数栈中弹出两个int变量v2和v1, v1是要进行位移的操作变量，
 * v2 指出要位移多少位
 * 1. int 变量只有32位，所以只取v2的前5个比特就足够表示位移数了
 * 2. Go 语言位操作符右侧必须是无符号整数，所以要对v2进行类型转换
 */
func (self *ISHL) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// long 算术右位移
type LSHR struct { base.NoOperandsInstruction }
func (self *LSHR) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//long 变量有64位 所以取v2的前6个比特
// Go 语言并没有java的 >>> 运算符， 为了达到无符号位移的目的，需要把
// v1 转换成无符号整数，位移操作之后，再转回有符号整数
// int 逻辑右位移
type IUSHR struct { base.NoOperandsInstruction }
func (self *IUSHR) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) * 0x1f
	result := int32(uint32(v1) >> s )
	stack.PushInt(result)
}


// int 算术右位移
type ISHR struct { base.NoOperandsInstruction }
func (self *ISHR) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) * 0x1f
	result := int32(uint32(v1) >> s )
	stack.PushInt(result)
}

// long 左位移
type LSHL struct { base.NoOperandsInstruction }
func (self *LSHL) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) * 0x1f
	result := int32(uint32(v1) >> s )
	stack.PushInt(result)
}

// long 逻辑右位移
type LUSHR struct { base.NoOperandsInstruction }
func (self *LUSHR) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) * 0x1f
	result := int32(uint32(v1) >> s )
	stack.PushInt(result)
}






