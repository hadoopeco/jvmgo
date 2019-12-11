package heap

import (
	"fmt"
	"strings"
)

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/8 23:58
 */

type MethodDescriptorParser struct {
	raw        string
	offset     int
	parsedDesc *MethodDescriptor
}

func parserMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (self *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	self.raw = descriptor
	fmt.Printf("parse = %v \n", self.raw)
	self.parsedDesc = &MethodDescriptor{}
	fmt.Printf("startParams raw = %v  offset = %v \n", self.raw, self.offset)
	self.startParams()
	fmt.Printf("parseParamTypes raw = %v  offset = %v \n", self.raw, self.offset)
	self.parseParamTypes() //解析参数
	self.endParams()
	self.parseReturnType() // 解析返回类型
	self.finish()
	return self.parsedDesc
}

func (self *MethodDescriptorParser) startParams() {
	if self.readUint8() != '(' {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		t := self.parseFieldType()
		fmt.Printf("parseParamTypes %v \n", t)
		if t != "" {
			self.parsedDesc.addParameterType(t)
		} else {
			break
		}
	}
}

//参数类型从字符转换成string
func (self *MethodDescriptorParser) parseFieldType() string {
	switch self.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.unreadUint8()
		return ""
	}
}

func (self *MethodDescriptorParser) parseObjectType() string {
	unread := self.raw[self.offset:]                 //先取出偏移量值到末尾的字符
	semicolonIndex := strings.IndexRune(unread, ';') //取剩余字符中的第一个字符';'的位置的index值
	if semicolonIndex == -1 {                        // 没找到则终止
		self.causePanic()
		return ""
	} else {
		objStart := self.offset - 1
		objEnd := self.offset + semicolonIndex + 1
		self.offset = objEnd                    // 偏移量移动到下一个位置
		descriptor := self.raw[objStart:objEnd] //读取开始到;位置的内容
		return descriptor
	}

}

func (self *MethodDescriptorParser) parseArrayType() string {
	arrStart := self.offset - 1
	self.parseFieldType()
	arrEnd := self.offset
	descriptor := self.raw[arrStart:arrEnd]
	return descriptor
}

func (self *MethodDescriptorParser) endParams() {
	if self.readUint8() != ')' {
		self.causePanic()
	}
}

//解析返回类型
func (self *MethodDescriptorParser) parseReturnType() {

	if self.readUint8() == 'V' {
		self.parsedDesc.returnType = "V" // void 返回类型
		return
	}

	self.unreadUint8()
	t := self.parseFieldType()
	if t != "" {
		self.parsedDesc.returnType = t
		return
	}

	self.causePanic()
}

func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.raw) {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + self.raw)
}

func (self *MethodDescriptorParser) readUint8() uint8 {

	b := self.raw[self.offset]
	self.offset++
	return b
}
func (self *MethodDescriptorParser) unreadUint8() {
	self.offset--
}
