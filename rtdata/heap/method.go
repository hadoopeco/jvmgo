package heap
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 10:12
 */
type Method struct {
	ClassMember
	maxStack		uint
	maxLocals		uint
	code 			[]byte
}
