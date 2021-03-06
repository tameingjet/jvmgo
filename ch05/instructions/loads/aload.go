package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ALOAD struct {
	base.Index8Instruction ;
}
func (self *ALOAD) Execute(frame *rtda.Frame)  {
	_aload(frame,uint(self.Index));
}
type ALOAD_0 struct {
	base.NoOperandsInstruction ;
}
func (self *ALOAD_0) Execute(frame *rtda.Frame)  {
	_iload(frame,uint(0));
}
type ALOAD_1 struct {
	base.NoOperandsInstruction ;
}
func (self *ALOAD_1) Execute(frame *rtda.Frame)  {
	_iload(frame,uint(1));
}
type ALOAD_2 struct {
	base.Index8Instruction ;
}
func (self *ALOAD_2) Execute(frame *rtda.Frame)  {
	_iload(frame,uint(2));
}
type ALOAD_3 struct {
	base.Index8Instruction ;
}
func (self *ALOAD_3) Execute(frame *rtda.Frame)  {
	_iload(frame,uint(2));
}
func _aload(frame *rtda.Frame,index uint)  {
	val := frame.LocalVars.GetRef(index);
	frame.OperandStack.PushRef(val) ;
}


