package loads

import "jvmgo/instructions/base"
import "jvmgo/rtdata"
/*
	加载指令从局部变量表获取变量, 然后推入操作数栈顶,
	加载指令共有33条, 按照操作变量的类型可以分为6类
	aload系列指令操作引用类型
	dload系列操作double类型变量
	fload系列操作float类型变量
	iload ...
	lload ...
 */
// Load reference from local variable
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rtdata.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rtdata.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rtdata.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rtdata.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rtdata.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtdata.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
