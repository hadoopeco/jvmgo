package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 9:46
 */
/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
//exception 是变长属性
 type ExceptionsAttribute struct {
 	exceptionIndexTable []uint16
 }

func (self *ExceptionsAttribute) readInfo(reader *ClassReader){
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable()[]uint16  {
	return self.exceptionIndexTable
}