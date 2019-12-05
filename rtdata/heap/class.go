package heap

/**
* Copyright (C) 2018
* All rights reserved
*
* @author: mark.wei
* @mail: marks@126.com
* Date: 2018/3/25 11:20

  方法区是运行时数据区的一块逻辑区域，由多个线程共享
  方法区主要存放class文件获取的信息，此外类变量也放在方法区里。
  当JVM第一次使用某个类时，它会搜索类路径，找到相应的class文件，然后读取并解析class文件，把相关信息放在方法区。

*/

import (
	"jvmgo/classfile"
)

type Class struct {
	accessFlag        uint16
	name              string //this class name
	superClassName    string //super class name
	interfaceNames    []string
	constantPool      *ConstantPool // 运行时常量指针
	fields            []*Field      //字段表
	methods           []*Method     // 方法表
	loader            *ClassLoader  // 类加载器指针
	superClass        *Class        // 超类指针
	interfaces        []*Class      // 接口指针
	instanceSlotCount uint          // 实例变量占据的空间大小
	staticSlotCount   uint          // 类变量占据的空间大小
	staticVars        Slots         // 静态变量
}

//把classFile结构体转换成 class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlag = cf.AccessFlags()
	class.superClassName = cf.SupperClassName()
	class.interfaceNames = cf.InterFaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

//判断是否为public class
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlag&ACC_PUBLIC
}
