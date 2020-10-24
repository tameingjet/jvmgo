package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
//算术右移：最高位填充符号位。正数填充0，负数填充1
//逻辑右移：最高位填充0
//左移都是补0
type ISHL struct {
	//int算术左位移
	base.NoOperandsInstruction;
}
//先从操作数栈中弹出两个int变量v2和v1。
//v1：要进行位移操作的变量
//v2：要移位多少比特
//位移之后把结果推入操作数栈

func (self *ISHL)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopInt();
	v1 := stack.PopInt() ;
	//int只有32位
	//只取v2的前5个比特就足够表示位移位数，而且go位移右侧数字必须无符号
	s := uint32(v2) & 0x1f ;
	result := v1 << s
	stack.PushInt(result) ;
}
type ISHR struct {
	//int算术右位移
	base.NoOperandsInstruction;
}
func (self *ISHR)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopInt();
	v1 := stack.PopInt() ;
	//int只有32位
	//只取v2的前5个比特就足够表示位移位数，而且go位移右侧数字必须无符号
	s := uint32(v2) & 0x1f ;
	result := v1 >> s
	stack.PushInt(result) ;
}

type IUSHR struct {
	//int逻辑右位移
	base.NoOperandsInstruction;
}
func (self *IUSHR)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopInt();
	v1 := stack.PopInt() ;
	s := uint32(v2) & 0x31 ;
	result := int32(v1 >> s)
	//先把v1转成无符号整数，位移操作后
	//再转回有符号
	stack.PushInt(result) ;
}
type LSHL struct {
	//long算术左位移
	base.NoOperandsInstruction;
}
func (self *LSHL)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopLong();
	v1 := stack.PopLong() ;
	//long64位
	//取v2的前6个比特
	s := uint32(v2) & 0x3f ;
	result := v1 >> s
	stack.PushLong(result) ;
}
type LSHR struct {
	//long算术左位移
	base.NoOperandsInstruction;
}
func (self *LSHR)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopLong();
	v1 := stack.PopLong() ;
	//long64位
	//取v2的前6个比特
	s := uint32(v2) & 0x3f ;
	result := v1 << s
	stack.PushLong(result) ;
}
type LUSHR struct {
	//long逻辑右位移
	base.NoOperandsInstruction;
}
func (self *LUSHR)Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack ;
	v2 := stack.PopLong();
	v1 := stack.PopLong() ;
	s := uint32(v2) & 0x31 ;
	result := int64(v1 >> s)
	//先把v1转成无符号整数，位移操作后
	//再转回有符号
	stack.PushLong(result) ;
}