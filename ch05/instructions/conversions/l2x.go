package conversions
import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type L2F struct {
	base.NoOperandsInstruction ;
}
func (self *L2F) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopLong() ;
	i:=float32(d)
	stack.PushFloat(i) ;

}
type L2D struct {
	base.NoOperandsInstruction ;
}

func (self *L2D) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopLong() ;
	i:=float64(d)
	stack.PushDouble(i) ;

}
type L2I struct {
	base.NoOperandsInstruction ;
}
func (self *L2I) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	d:=stack.PopLong() ;
	i:=int32(d)
	stack.PushInt(i) ;
}