package heap

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/29 0:29
 */

// keywords table
const (
	ACC_PUBLIC       = 0x0001 //class  field method  public
	ACC_PRIVATE      = 0x0002 // 	    field method  private
	ACC_PROTECTED    = 0x0004 //       field method  protected
	ACC_STATIC       = 0x0008 //       field method  static
	ACC_FINAL        = 0x0010 // class field method  final
	ACC_SUPER        = 0x0020 // class               super
	ACC_SYNCHRONIZED = 0x0020 //             method  synchronized
	ACC_VOLATILE     = 0x0040 //       field         volatile
	ACC_BRIDGE       = 0x0040 //             method  bridge
	ACC_TRANSIENT    = 0x0080 //       field         transient
	ACC_VARARGS      = 0x0080 //             method  varargs
	ACC_NATIVE       = 0x0100 //  			  method  native
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class       method
	ACC_STRICT       = 0x0800 //             method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field

)
