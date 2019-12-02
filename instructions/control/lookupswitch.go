package control

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
 * Date: 2018/3/27 22:34
 *
 * java语言中的switch-case语句有两种实现方式，如果case值可以编码成一个索引表，则实现tableswitch指令
 * 否则实现成lookuptable 指令
 *
 * lookupswitch指令
 *
 * int chooseNear(i){
 *    switch(i){
 *     case -100:  return -1;
 *     case 0:  return 0;
 *     case 100:  return 1;
 *	   default: return -1;
 *	}
 * }
 *
 */


type LOOKUP_SWITCH struct {
 	defaultOffset   int32
 	npairs          int32
 	matchOffsets    []int32
 }


 func(self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
 	reader.SkipPadding()
 	self.defaultOffset = reader.ReadInt32()
 	self.npairs = reader.ReadInt32()
 	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
 }

 /*
 	matchOffsets 类似于Map, key是case的值, value是偏移量
 	Execute 方法先从操作数栈中弹出一个int值, 然后再用它查找matchOffsets的key
 	如果能找到,则按照value给出的偏移量跳转,不能则按照default值跳转
  */
 func (self *LOOKUP_SWITCH) Execute(frame *rtdata.Frame) {
 	 key := frame.OperandStack().PopInt()
 	 for i := int32(0); i < self.npairs*2 ; i += 2 {
		if self.matchOffsets[i]  == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	 }
	 base.Branch(frame, int(self.defaultOffset))
 }