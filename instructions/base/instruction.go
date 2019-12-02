package base

import "jvmgo/rtdata"

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/27 17:40
 */

type Instruction interface {
	FetchOperands(reader *BytecodeReader)//从字节码中提取操作数
	Execute(frame *rtdata.Frame) //执行指令逻辑
}

type NoOperandsInstruction struct {
// empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

//跳转指令 offset 存放跳转偏移量
type BranchInstruction struct {	Offset int}

//从字节中读取一个uint16整数  给offset赋值
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}
//存储和加载类指令需要根据索引存取局部变量表, 索引由单字节操作数给出, 把这类指令抽象成Index8Instruction
//index表示局部变量表索引
type Index8Instruction struct {	Index uint}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

//有一些指令需要访问运行时常量池,常量池索引由双字节操作数给出
//index表示常量池索引
type Index16Instruction struct { Index uint}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}



