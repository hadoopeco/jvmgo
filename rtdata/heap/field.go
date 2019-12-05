package heap

import "jvmgo/classfile"

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 10:43
 */
type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

//把classfile.MemberInfo 信息拷贝到field
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}

	return fields
}

// prepare for classloader.allocAndInitStaticVars.cpIndex
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Field) SlotId() uint {
	return self.slotId
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) Descriptor() string {
	return self.descriptor
}
