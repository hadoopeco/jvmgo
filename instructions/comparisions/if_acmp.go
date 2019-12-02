package comparisions

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
 * Date: 2018/3/27 22:14
 */
 // Branch if reference comparison succeeds

 type IF_ACMPEQ struct { base.BranchInstruction }
 func (self *IF_ACMPEQ) Execute(frame *rtdata.Frame) {
	if _acmp(frame) {
		 base.Branch(frame, self.Offset)
	 }
 }

 type IF_ACMPNE struct { base.BranchInstruction }

 func (self *IF_ACMPNE) Execute(frame *rtdata.Frame) {
	 if !_acmp(frame) {
		 base.Branch(frame, self.Offset)
	 }
 }


 func _acmp(frame *rtdata.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
 }