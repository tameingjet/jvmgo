package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

//求余
type IREM struct {
	base.NoOperandsInstruction ;
}
//先从操作数栈中弹出两个int变量，求余，然后把结果推入操作数栈
//对int或long变量做除法和求余运算时，是有可能抛出ArithmeticException异常
func (self *IREM) Execute(frame rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopInt() ;
	v1 := stack.PopInt() ;
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/by zero") ;
	}
	result := v1%v2 ;
	stack.PushInt(result);
}
type FREM struct {
	base.NoOperandsInstruction ;
}
//浮点数有Infinity（无穷大）值，所以即使除0也没异常抛出
func (self *FREM) Execute(frame rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopFloat() ;
	v1 := stack.PopFloat() ;
	result := float32(math.Mod(float64(v1), float64(v2))) ;
	stack.PushFloat(result);
}
type DREM struct {
	base.NoOperandsInstruction ;
}
func (self *DREM) Execute(frame rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopDouble() ;
	v1 := stack.PopDouble() ;
	result := math.Mod(v1,v2) ;
	stack.PushDouble(result);
}
type LREM struct {
	base.NoOperandsInstruction ;
}
func (self *LREM) Execute(frame rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopLong() ;
	v1 := stack.PopLong() ;
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/by zero") ;
	}
	result := v1%v2 ;
	stack.PushLong(result);
}
