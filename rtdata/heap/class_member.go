package heap

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/25 10:34
 */

type ClassMember struct {
	accessFlags			uint16
	name				string
	descriptor			string
	class				*Class
}
