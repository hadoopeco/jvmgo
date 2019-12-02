package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/27 22:16
 */


 type GOTO struct { base.BranchInstruction }

 func (self *GOTO) Execute(frame *rtdata.Frame){
	base.Branch(frame, self.Offset)
 }
