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
 * Date: 2018/3/27 22:09
 */

//if int1 == int2
type IF_ICMPEQ struct{	base.BranchInstruction }
func (self *IF_ICMPEQ) Execute(frame *rtdata.Frame){

	if val1, val2 := _icmp(frame);val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

// if int1 != int2
type IF_ICMPNE struct{	base.BranchInstruction }
func (self *IF_ICMPNE) Execute(frame *rtdata.Frame){
	if val1, val2 := _icmp(frame); val1 != val2 {
		base.Branch(frame, self.Offset)
	}
 }


 type IF_ICMPLT struct{	base.BranchInstruction }
 func (self *IF_ICMPLT) Execute(frame *rtdata.Frame){
	 if val1, val2 := _icmp(frame);val1 < val2 {
		base.Branch(frame, self.Offset)
	}
 }

 type IF_ICMPLE struct{	base.BranchInstruction }
 func (self *IF_ICMPLE) Execute(frame *rtdata.Frame){
	 if val1, val2 := _icmp(frame);val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
 }


 type IF_ICMPGT struct{	base.BranchInstruction }
 func (self *IF_ICMPGT) Execute(frame *rtdata.Frame){
	 if val1, val2 := _icmp(frame);val1 > val2 {
		base.Branch(frame, self.Offset)
	}
 }

 type IF_ICMPGE struct{	base.BranchInstruction }
 func (self *IF_ICMPGE) Execute(frame *rtdata.Frame){
	 if val1, val2 := _icmp(frame);val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
 }

 func _icmp(frame *rtdata.Frame) (v1, v2 int32) {
 	stack := frame.OperandStack()
 	v2 = stack.PopInt()
 	v1 = stack.PopInt()
 	return v1, v2
 }












