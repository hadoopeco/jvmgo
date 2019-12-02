package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtdata"
	"fmt"
)

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/27 17:56
 */

 // Do nothing
 type NOP struct { 	base.NoOperandsInstruction }

 func (self *NOP) Execute(frame *rtdata.Frame){
	fmt.Println("tsts")
 }