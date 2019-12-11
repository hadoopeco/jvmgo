package references

import (
	"fmt"
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
 * Date: 2019/12/8 0:07
 */

//invoke static method
type INVOKE_STATIC struct{ base.Index16Instruction }

/*
 如果解析符号引用后得到方法M, 那么M必须是静态方法,且M不能是类初始化方法
 类初始化方法只能被java虚拟机调用,不能通过invokestatic指令调用
 对于invokestatic指令, M就是要最终执行的方法,调用InvokeMethod执行该方法
*/
func (self *INVOKE_STATIC) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	fmt.Printf("ref = %v, method === %v \n", methodRef, resolvedMethod)
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod) //方法执行完后执行 return指令
}
