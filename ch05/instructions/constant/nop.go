package constant

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//nop指令是最简单的一条指令，因为它什么也不做
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	//什么都不用做
}
