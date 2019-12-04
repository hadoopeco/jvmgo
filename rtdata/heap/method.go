package heap

import "jvmgo/classfile"

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
	maxStack  uint   //操作数栈的大小
	maxLocals uint   //局部变量表的大小
	code      []byte //存放方法字节码
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))

	for i, memberInfo := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(memberInfo)
		methods[i].copyAttributes(memberInfo)
	}

	return methods
}

func (self Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}

}
