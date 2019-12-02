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
 * Date: 2018/3/27 17:58
 */

/*
	常量指令把常量推入操作数栈顶, 常量可以来自三个地方:隐含在操作码里, 操作数和运行时常量值
	常量指令共有21条
 */

type ACONST_NULL struct { base.NoOperandsInstruction }
func (self *ACONST_NULL) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushRef(nil)
}


type DCONST_0 struct {	base.NoOperandsInstruction }
func (self *DCONST_0) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushDouble(0.0)
}

type ICONST_M1 struct {	 base.NoOperandsInstruction }
func (self *ICONST_M1) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(-1)
}

type DCONST_1 struct {	base.NoOperandsInstruction }
func (self *DCONST_1) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {	base.NoOperandsInstruction }
func (self *FCONST_0) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {	base.NoOperandsInstruction }
func (self *FCONST_1) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {	base.NoOperandsInstruction }
func (self *FCONST_2) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_0 struct {	base.NoOperandsInstruction }
func (self *ICONST_0) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {	base.NoOperandsInstruction }
func (self *ICONST_1) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {	base.NoOperandsInstruction }
func (self *ICONST_2) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {	base.NoOperandsInstruction }
func (self *ICONST_3) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {	base.NoOperandsInstruction }
func (self *ICONST_4) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {	base.NoOperandsInstruction }
func (self *ICONST_5) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushInt(5)
}

type LCONST_1 struct {	base.NoOperandsInstruction }
func (self *LCONST_1) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushLong(1)
}

type LCONST_0 struct {	base.NoOperandsInstruction }
func (self *LCONST_0) Execute(frame *rtdata.Frame){
	frame.OperandStack().PushLong(0)
}