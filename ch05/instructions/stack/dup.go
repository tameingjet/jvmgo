package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
//dup指令复制栈顶的单个变量

type DUP struct {
	/*
	bottom -> top
	[...][c][b][a]
	             \_
	               |
	               V
	[...][c][b][a][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot := stack.PopSlot();
	stack.PushSlot(slot) ;
	stack.PushSlot(slot);
}
type DUP_X1 struct {
	/*
	bottom -> top
	[...][c][b][a]
	          __/
	         |
	         V
	[...][c][a][b][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP_X1) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot1 := stack.PopSlot();
	slot2 := stack.PopSlot() ;
	stack.PushSlot(slot1) ;
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1) ;
}

type DUP_X2 struct {
	/*
	bottom -> top
	[...][c][b][a]
	       _____/
	      |
	      V
	[...][a][c][b][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP_X2) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot1 := stack.PopSlot();
	slot2 := stack.PopSlot() ;
	slot3 := stack.PopSlot() ;
	stack.PushSlot(slot1) ;
	stack.PushSlot(slot3)
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1) ;
}
type DUP2 struct {
	/*
	bottom -> top
	[...][c][b][a]____
	          \____   |
	               |  |
	               V  V
	[...][c][b][a][b][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP2) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot1 := stack.PopSlot();
	slot2 := stack.PopSlot() ;
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1)
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1) ;
}
type DUP2_X1 struct {
	/*
	bottom -> top
	[...][c][b][a]
	       _/ __/
	      |  |
	      V  V
	[...][b][a][c][b][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP2_X1) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot1 := stack.PopSlot();
	slot2 := stack.PopSlot() ;
	slot3 := stack.PopSlot() ;
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1) ;
}
type DUP2_X2 struct {
	/*
	bottom -> top
	[...][d][c][b][a]
	       ____/ __/
	      |   __/
	      V  V
	[...][b][a][d][c][b][a]
	*/
	base.NoOperandsInstruction ;
}
func (self *DUP2_X2) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack ;
	slot1 := stack.PopSlot();
	slot2 := stack.PopSlot() ;
	slot3 := stack.PopSlot() ;
	slot4 :=  stack.PopSlot() ;
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2) ;
	stack.PushSlot(slot1) ;
}