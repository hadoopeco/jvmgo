package heap

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 10:30
 */

import "jvmgo/classpath"

type ClassLoader struct {
	cp             *classpath.Classpath
	classMap       map[string]*Class   //loaded classes

}
