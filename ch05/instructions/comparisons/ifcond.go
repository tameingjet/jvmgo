package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
//if<cond>指令把操作数栈顶的int变量弹出，然后跟0进行比较，满足条件
//则跳转
type IFEQ struct {
	//ifeq：x==0
	base.BranchInstruction ;
}

func (self *IFEQ)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val == 0 {
		base.Branch(frame,self.Offset)
	}
}
type IFNE struct {
	//x!=0
	base.BranchInstruction ;
}
func (self *IFNE)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val != 0 {
		base.Branch(frame,self.Offset)
	}
}
type IFLT struct {
	//x<0
	base.BranchInstruction ;
}
func (self *IFLT)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val < 0 {
		base.Branch(frame,self.Offset)
	}
}
type IFLE struct {
	//x<=0
	base.BranchInstruction ;
}
func (self *IFLE)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val <= 0 {
		base.Branch(frame,self.Offset)
	}
}
type IFGT struct {
	//x>0
	base.BranchInstruction ;
}
func (self *IFGT)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val > 0 {
		base.Branch(frame,self.Offset)
	}
}
type IFGE struct {
	//x>=0
	base.BranchInstruction ;
}
func (self *IFGE)Execute(frame *rtda.Frame)  {
	val := frame.OperandStack.PopInt() ;
	if val >= 0 {
		base.Branch(frame,self.Offset)
	}
}