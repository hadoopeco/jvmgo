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
 * Date: 2019/12/5 21:19
 */

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.InCompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1) //error missing -1
	if ref == nil {
		if methodRef.Name() == "println" {
			_printf(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}

	//保证protected 类只能被该类或子类方法调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		resolvedMethod.Class().IsSupperClassOf(currentClass) &&
		ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	//调用超类中非构造函数的函数，且当前ACC_SUPPER标志被设置, 需要查找最终要调用的方法
	//若非如此,resolvedMethod 就是要调用的方法
	//methodTobeInvoke := resolvedMethod
	//methodTobeInvoke := resolvedMethod
	//if currentClass.IsSuper() && resolvedMethod.Class().IsSupperClassOf(currentClass) &&
	//	resolvedMethod.Name() != "<init>"{
	//
	//}
	methodTobeInvoke := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodTobeInvoke == nil || methodTobeInvoke.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodTobeInvoke)
}

func _printf(stack *rtdata.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("int : %v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
