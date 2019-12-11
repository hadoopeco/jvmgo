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
 * Date: 2019/12/5 21:17
 */

// Invoke instance method, special handling for superclass, private, and instance initialization
// method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef) //解析当前类的符号引用
	resolvedMethod := methodRef.ResolvedMethod()
	resolvedClass := methodRef.ResolvedClass()

	//如果从方法符号引用中解析出来的类是C, 方法是M, 如果M是构造函数,则声明M的类必须是C,否则抛出NosuchMethodError异常
	if resolvedMethod.Class() != resolvedClass && resolvedMethod.Name() == "<init>" {
		panic("java.lang.NoSuchMethodError")
	}

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 返回距离操作数栈顶n个单元的引用变量
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1) //todo:error code
	if ref == nil {
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
	methodTobeInvoked := resolvedMethod //resolvedClass write wrong
	if currentClass.IsSuper() && resolvedClass.IsSupperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodTobeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodTobeInvoked == nil || methodTobeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodTobeInvoked)
}
