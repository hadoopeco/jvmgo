package heap

import (
	"jvmgo/classfile"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 10:12
 */
type Method struct {
	ClassMember
	maxStack     uint   //操作数栈的大小
	maxLocals    uint   //局部变量表的大小
	argSlotCount uint   //参数个数
	code         []byte //存放方法字节码
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))

	for i, memberInfo := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(memberInfo)
		methods[i].copyAttributes(memberInfo)
		methods[i].calcArgSlotCount()
	}

	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}

}

func (self *Method) calcArgSlotCount() {
	parseDescriptor := parserMethodDescriptor(self.descriptor)
	for _, paramType := range parseDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}

	if !self.IsStatic() {
		self.argSlotCount++ //编译器默认添加 this 作为参数
	}

}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

// getters
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
