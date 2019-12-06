package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
)

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/6 10:25
 */

type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtdata.Frame) {
	frame.Thread().PopFrame()
}

// Return reference from method
//type ARETURN struct{ base.NoOperandsInstruction }
//
//func (self *ARETURN) Execute(frame *rtdata.Frame) {
//	thread := frame.Thread()
//	currentFrame := thread.PopFrame()
//	invokerFrame := thread.TopFrame()
//	ref := currentFrame.OperandStack().PopRef()
//	invokerFrame.OperandStack().PushRef(ref)
//}
