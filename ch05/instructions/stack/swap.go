package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
//swap指令交换栈顶的两个变量
type SWAP struct {
	base.NoOperandsInstruction ;
}

func (self *SWAP) Execute(frame rtda.Frame)  {
	stack := frame.OperandStack ;
	slot1:=stack.PopSlot() ;
	slot2:=stack.PopSlot() ;
	stack.PushSlot(slot1) ;
	stack.PushSlot(slot2) ;
}
