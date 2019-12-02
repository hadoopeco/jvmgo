package classfile

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/24 9:21
 */


 //标记属性 无数据
type MarkerAttribute struct {

}
func (self *MarkerAttribute) readInfo(cp *ClassReader) {
	//doing nothing
}

/*
Deprecated_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

