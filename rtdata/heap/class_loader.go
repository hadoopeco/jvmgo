package heap

/**
* Copyright (C) 2018
* All rights reserved
*
* @author: mark.wei
* @mail: marks@126.com
* Date: 2018/3/25 10:30



 ClassLoader 依赖ClassPath 来搜索和读取class文件, cp字段保存Classpath指针
 classMap记录已经加载的类数据, key是类名,
*/

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"strings"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class //loaded classes

}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return self.loadNonArrayClass(name)
}

/*
 先查找classMap, 看类是否被加载, 如果有直接返回
 数组类和普通类不同,它的数据不来自于class文件,而是有虚拟机在运行时生成
*/
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	fmt.Println("class name :" + name)
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Load %s from %s]\n", name, entry)
	}
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException : " + name)
	}
	return data, entry
}

//类的加载分三步骤 首先把class文件数据读入内存;然后解析class文件,生成虚拟机可使用的类数据,并方入方法区;最后进行链接
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

//把class文件数据转换成class结构体
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

//所有类只有一个父类,除Object外,递归调用loadClass加载父类
func resolveSuperClass(class *Class) {
	if strings.Index(class.name, "java/lang/Object") < 0 {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}

}

//递归调用loadClass 加载interface
func resolveInterfaces(class *Class) {
	interfacesCount := len(class.interfaceNames)
	if interfacesCount > 0 {
		class.interfaces = make([]*Class, interfacesCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo:
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 计算实例字段的个数, 并给他们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			_calcSlotId(field, &slotId)
		}
	}
	class.instanceSlotCount = slotId
}

//计算静态字段的个数, 并且给他们编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			_calcSlotId(field, &slotId)

		}
	}
	class.staticSlotCount = slotId
}

func _calcSlotId(field *Field, slotId *uint) {
	field.slotId = *slotId
	*slotId++
	if field.isLongOrDouble() {
		*slotId++
	}
}

// 给类变量分配空间, 然后赋初始值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	soltId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(soltId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(soltId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(soltId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(cpIndex, val)
		case "Ljava.lang.String;":
			panic("todo")
		}
	}
}
