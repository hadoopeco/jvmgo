package loads

import "jvmgo/instructions/base"
import "jvmgo/rtdata"

// Load long from local variable
type LLOAD struct { base.Index8Instruction }
func (self *LLOAD) Execute(frame *rtdata.Frame){
	_lload(frame, 1)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }
func (self *LLOAD_0) Execute(frame *rtdata.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtdata.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtdata.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }
func (self *LLOAD_3) Execute(frame *rtdata.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
