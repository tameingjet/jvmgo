package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//栈指令直接对操作数栈进行操作，共9条：pop和pop2指令将栈
//顶变量弹出，dup系列指令复制栈顶变量，swap指令交换栈顶的两
//个变量。
type POP struct {
	//pop指令只能用于弹出int，float等占用一个操作数栈位置的变量
	//double和long用pop2
	base.NoOperandsInstruction ;
}

func (self *POP) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	stack.PopSlot()
}
type POP2 struct {
	//弹出double，long等占两个位置的
	base.NoOperandsInstruction ;
}
func (self *POP2) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	stack.PopSlot()
	stack.PopSlot()
}
