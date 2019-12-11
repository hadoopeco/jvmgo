package base

import (
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
)

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/8 0:18
 */

func InvokeMethod(invokeFrame *rtdata.Frame, method *heap.Method) {
	//创建新的frame 推入thread
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	//将需要的参数,从调用方拷贝到执行的方法
	argsCount := int(method.ArgSlotCount())
	if argsCount > 0 {
		for i := argsCount - 1; i >= 0; i-- {
			slot := invokeFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic("error")
		}
	}

}
