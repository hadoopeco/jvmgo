package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 0:32
 */

 type ConstantClassInfo struct {
 	cp ConstantPool
 	nameIndex uint16
 }

 func (self *ConstantClassInfo) readInfo (reader *ClassReader){
 	self.nameIndex = reader.readUint16()
 }

func (self *ConstantClassInfo) String(reader *ClassReader) string {
	return self.cp.getUtf8(self.nameIndex)
}