package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 9:32
 */
/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
 type ConstantValueAttribute struct {
 	constantValueIndex uint16
 }

 func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
 }

 func (self *ConstantValueAttribute) ConstantValueIndex(reader *ClassReader) uint16{
 	return self.constantValueIndex
 }