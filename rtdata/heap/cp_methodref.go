package heap

import "jvmgo/classfile"

/**
 * Copyright (C) 2019
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: wbmark@gmail.com
 * Date: 2019/12/4 11:07
 */

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
