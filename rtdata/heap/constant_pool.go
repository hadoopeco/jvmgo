package heap
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:14
 */

type Constant interface{}


type ConstantPool struct {
	class 		*Class
	constants  	[]Constant
}