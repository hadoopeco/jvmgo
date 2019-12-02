package classfile
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/20 19:08
 */
import "fmt"


/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields [] *MemberInfo
	methods [] *MemberInfo
    attributes [] AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error)  {
	defer func() {
		if r := recover(); r != nil{
			var ok  bool
			_, ok = r.(error)

			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{ classData }
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader)   {
	self.readCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// reader the file header to make sure the file is java Class file
func (self *ClassFile) readCheckMagic(reader *ClassReader)  {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("Class file format error ")
	}
}


func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) ClassName() string{
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SupperClassName() string  {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "Object no superClass"  // java.lang.Object 没有超类
}

func (self *ClassFile) InterFaceNames() []string{
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader)  {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 46, 47,48,50,51,52:
		if self.minorVersion == 0{
			return
		}
	}

	panic("java.lang.UnSupportClassVersion!")
}

func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Methods() []*MemberInfo{
	return self.methods
}

func (self *ClassFile) Fields() []*MemberInfo{
	return self.fields
}
