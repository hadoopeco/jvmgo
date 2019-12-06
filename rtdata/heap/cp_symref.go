package heap

/**
* Copyright (C) 2019
* All rights reserved
*
* @author: mark.wei
* @mail: wbmark@gmail.com
* Date: 2019/12/3 13:18


 cp字段存放符号引用所在的常量池指针,这样可以通过符号引用访问到运行时常量池,进一步可以访问到类数据
 className 字段存放类名, class字段缓存解析后的类结构体指针,这样类符号引用只需要解析一次就可以了,后续可以直接使用缓存值
 对于类符号引用, 只要有类名,就可以解析符号引用
 对于字段,首先要解析类符号引用得到数据,然后用字段名和描述符查找字段数据
 方法符号的解析过程和字段符号的解析过程类似

						SymRef
						  |
					______|________
                   |               |
                ClassRef        MemberRef
                    _______________|______________
                   |               |              |
           InterfaceMethodRef   MethodRef      FieldRef
*/
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	//如果类d想访问c, 需要满足两个条件之一: c是public, 或者c 和d 在同一个包内
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
