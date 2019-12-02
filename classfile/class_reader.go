package classfile
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/20 17:20
 */
import "encoding/binary"

type ClassReader struct {
	data []byte
}

//8比特有符号整数
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//16比特无符号整数
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//32比特有符号整数
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
//64比特无符号整数
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// read uint16 table, table's size start point base on header
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}

	return s
}

// reader any length of identified bytes
func (self *ClassReader) readBytes(n uint32) []byte{
	bytes := self.data[:n]
	self.data = self.data[n:]

	return bytes
}
