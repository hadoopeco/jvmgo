package loads

import "jvmgo/instructions/base"
import "jvmgo/rtdata"

// Load double from local variable
type DLOAD struct{ base.Index8Instruction }

func (self *DLOAD) Execute(frame *rtdata.Frame) {
	_dload(frame, self.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtdata.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtdata.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtdata.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtdata.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
