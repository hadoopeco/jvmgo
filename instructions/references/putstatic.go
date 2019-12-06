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
 * Date: 2019/12/5 16:55
 */
// 静态变量赋值
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtdata.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstants(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	fmt.Printf("PUT_STATIC = %v \n", field)
	class := field.Class()

	//解析出的字段如不是静态变量则报错
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//解析出的静态变量若是final修饰,则只能在类初始化时赋值,否则报错   // Method java/lang/Object."<init>"
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	//根据字段类型从操作数中弹出相应的值，给静态变量赋值
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		//todo:
	}
}
