package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
	"math"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/27 19:04
 */
 /**
 * * 算术指令又可以进一步分为加法指令(add)指令，减法(sub)指令,乘法(mul)指令
 * 	除法(div)指令 求余(rem)指令 和 取反(neg)指令
  */


// 先从操作数栈中弹出两个int变量，求余，然后再把结果push进操作数栈， 对double 或long变量做除法和求余运算
// 有可能抛出ArithmeticException 异常
type IREM struct{	base.NoOperandsInstruction }
func (self *IREM) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic(" java.lang.ArithmeticException : / by zero ")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}



// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
type DREM struct{	base.NoOperandsInstruction }
// 先从操作数栈中弹出两个int变量，求余，然后再把结果push进操作数栈
func (self *DREM) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2 )
	stack.PushDouble(result)
}