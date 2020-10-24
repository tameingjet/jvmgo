package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type I2F struct {
	base.NoOperandsInstruction ;
}
func (self *I2F) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=float32(d)
	stack.PushFloat(i) ;

}
type I2D struct {
	base.NoOperandsInstruction ;
}

func (self *I2D) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=float64(d)
	stack.PushDouble(i) ;

}
type I2L struct {
	base.NoOperandsInstruction ;
}
func (self *I2L) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=int64(d)
	stack.PushLong(i) ;
}
// Convert int to byte
type I2B struct {
	base.NoOperandsInstruction ;
}
func (self *I2B) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=int32(int8(d))
	stack.PushInt(i) ;
}
// Convert int to char
type I2C struct {
	base.NoOperandsInstruction ;
}

func (self *I2C) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=int32(uint16(d))
	stack.PushInt(i) ;

}
// Convert int to short
type I2S struct {
	base.NoOperandsInstruction ;
}
func (self *I2S) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopInt() ;
	i:=int64(d)
	stack.PushLong(i) ;
}