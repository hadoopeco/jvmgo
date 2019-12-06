package heap

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/5 20:32
 */

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

// self extends c
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
