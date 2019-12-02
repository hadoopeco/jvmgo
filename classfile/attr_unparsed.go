package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 9:15
/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}
/**
*jvm defined the attributes
* ConstantValue  	field_info
* Code              member_info
* Exceptions        method_info
* SourceFile        ClassFile
*/
func (self *UnparsedAttribute) readInfo(read *ClassReader)  {
	self.info = read.readBytes(self.length)
}

