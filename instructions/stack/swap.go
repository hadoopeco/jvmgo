package stack

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
 * Date: 2018/3/27 18:48
 */

type SWAP struct{
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]

 Swap the top two operand stack values
 swap 指令交换栈顶的两个变量
*/
func (self *SWAP) Execute(frame *rtdata.Frame)  {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}


