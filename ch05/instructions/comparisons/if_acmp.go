package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IF_ACMPEQ struct {
	//ifeq：x==0
	base.BranchInstruction ;
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack;
	ref2 := stack.Popref();
	ref1:=stack.Popref();
	if ref1 == ref2 {
		base.Branch(frame,self.Offset);
	}
}
type IF_ACMPNE struct {
	//ifne：x!=0
	base.BranchInstruction ;
}
func (self *IF_ACMPNE) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack;
	ref2 := stack.Popref();
	ref1:=stack.Popref();
	if ref1 != ref2 {
		base.Branch(frame,self.Offset);
	}
}