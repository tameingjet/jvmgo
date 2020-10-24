package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//if_icmp<cond>指令把栈顶的两个int变量弹出，然后进行比较，
//满足条件则跳转
type IF_ICMPEQ struct {
	//ifeq：x==0
	base.BranchInstruction ;
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1==val2 {
		base.Branch(frame,self.Offset);
	}

}
type IF_ICMPNE struct {
	//ifNE：x!=0
	base.BranchInstruction ;
}
func (self *IF_ICMPNE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1!=val2 {
		base.Branch(frame,self.Offset);
	}

}
type IF_ICMPLT struct {
	//iflt：x<0
	base.BranchInstruction ;
}
func (self *IF_ICMPLT) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1<val2 {
		base.Branch(frame,self.Offset);
	}

}
type IF_ICMPLE struct {
	//ifle：x<=0
	base.BranchInstruction ;
}
func (self *IF_ICMPLE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1<=val2 {
		base.Branch(frame,self.Offset);
	}

}
type IF_ICMPGT struct {
	//ifgt：x>0
	base.BranchInstruction ;
}
func (self *IF_ICMPGT) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1>val2 {
		base.Branch(frame,self.Offset);
	}

}
type IF_ICMPGE struct {
	//ifge：x>=0
	base.BranchInstruction ;
}
func (self *IF_ICMPGE) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack;
	val2:= stack.PopInt();
	val1:= stack.PopInt();
	if val1>=val2 {
		base.Branch(frame,self.Offset);
	}

}

