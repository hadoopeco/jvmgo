package rtdata

import "jvmgo/rtdata/heap"

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 11:19
 */

type Slot struct {
	num int32
	ref *heap.Object
}