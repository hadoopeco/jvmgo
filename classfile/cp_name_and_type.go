package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 1:02
/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
 type ConstantNameAndTypeInfo struct {
 	nameIndex 		uint16
 	descriptorIndex uint16
 }

 func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
 	self.nameIndex = reader.readUint16()
 	self.descriptorIndex = reader.readUint16()
 }