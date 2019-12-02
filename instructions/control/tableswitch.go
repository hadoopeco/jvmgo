package control

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
 * Date: 2018/3/27 22:18

 * java语言中的switch-case语句有两种实现方式，如果case值可以编码成一个索引表，则实现tableswitch指令
 * 否则实现成lookuptable 指令
 *
 * tableswitch指令
 *
 * int chooseNear(i){
 *    switch(i){
 *     case 0:  return 0;
 *     case 1:  return 1;
 *     case 2:  return 2;
 *     case 3:  return 3;
 *     case 4:  return 4;
 *	  default: return -1;
 *	}
 * }
 *
 *  tableswitch
 *  <0-3 byte pad>
 *  defaultbyte1
 *  defaultbyte2
 *  defaultbyte3
 *  defaultbyte4
 *  lowbyte1
 *  lowbyte2
 *  lowbyte3
 *  lowbyte4
 *  highbyte1
 *  highbyte2
 *  highbyte3
 *  highbyte4
 *  jump offsets...
*/
 //Access jump table by index and jump
 type TABLE_SWITCH struct {
	defaultOffset int32
	low 		  int32
	high	      int32
	jumpOffsets   []int32
 }

/*
  tableswitch 指令操作码后面有0 - 3个字节的padding, 以保证defaultOffset在字节码中的地址是4的倍数
  defaultOffset 对应于默认情况下跳转的偏移量, low 和high 记录case的取值范围
  jumpOffsets  是一个索引表, 里面存放high-low+1 个值, 对应于各种case情况下跳转所需的字节码偏移量
*/
 func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
 }



/*
 先从操作数中弹出一个int值，看是否在low和high的取值范围,如果在,就按从jumpOffsets中查到的偏移量跳转
 如果不在,就按照default跳转
*/
 func (self *TABLE_SWITCH) Execute(frame *rtdata.Frame){
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index - self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
 }
