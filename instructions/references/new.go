package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
)

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/5 16:16
 */

// Create new object
type New struct{ base.Index16Instruction }

func (self *New) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classInfo := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classInfo.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
