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

	ConstantPool ---1-------1---- Class          ClassMember
										|1               |
										|          ______|______
										|		   |           |
										|-----*- Field      Method
										|					  |*
										|_____________________|
*/

import (
	"jvmgo/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
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
	initStarted       bool          //标识<clinit>方法是否已经开始运行
}

//把classFile结构体转换成 class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SupperClassName()
	class.interfaceNames = cf.InterFaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

//判断是否为public class
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) Name() string {
	return self.name
}
func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

// jvms 5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

//获取Main方法
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

//获取构造方法
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}
