package stack

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
 * Date: 2018/3/27 18:37
 */



type POP struct { 	base.NoOperandsInstruction }
/*
 把操作数栈顶的变量弹出,用于弹出int, float
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 把栈顶的变量弹出,用于弹出double 和long，此类型变量在操作数栈中占用两个位置
type POP2 struct {	base.NoOperandsInstruction }
/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (self *POP2) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
