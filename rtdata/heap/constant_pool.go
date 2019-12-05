package heap

import (
	"fmt"
	"jvmgo/classfile"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:14
 */

/*
	运行时常量主要放两类信息 字面量(literal) 和 符号引用(symbolic reference)
	字面量包括整数,浮点数和字符串字面量;
	符号引用包括类符号引用,字段符号引用,方法符号引用和接口方法符号引用
	ConstantPool ---1-------1---- Class          ClassMember
									|1               |
									|          ______|______
									|		   |           |
									|-----*- Field      Method
									|					  |*
									|_____________________|
*/
type Constant interface{}

type ConstantPool struct {
	class     *Class
	constants []Constant
}

// 根据索引获取常量值
func (self ConstantPool) getConstant(index uint) Constant {
	if c := self.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No Constant at index %d", index))
}

//把classfile 里的常量池转换成运行时常量池
func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfcp)
	constants := make([]Constant, cpCount)

	//运行时常量池
	rtcp := &ConstantPool{class, constants}
	//The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount; i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			constants[i] = intInfo.Value() // int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			constants[i] = floatInfo.Value() // float32
		// long 和 double 在常量池里占2个位置
		case *classfile.ConstantDoubleInfo:
			dInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			constants[i] = dInfo.Value() // float64
			i++
		case *classfile.ConstantLongInfo:
			lInfo := cpInfo.(*classfile.ConstantLongInfo)
			constants[i] = lInfo.Value() // int64
			i++
		case *classfile.ConstantStringInfo:
			sInfo := cpInfo.(*classfile.ConstantStringInfo)
			constants[i] = sInfo.String() //string

		default:
			//
		}
	}

	return rtcp
}
