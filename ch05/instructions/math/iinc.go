package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//iinc指令给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint ;//局部变量表索引
	Const int32 ;//常量值
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader)  {
	//从字节码里读取操作数
	self.Index = uint(reader.ReadUint8()) ;
	self.Const = int32(reader.ReadInt8()) ;
}
func (self *IINC) Execute(frame *rtda.Frame)  {
	localVars:=frame.LocalVars ;
	val := localVars.GetInt(self.Index) ;
	val+=self.Const ;
	localVars.SetInt(self.Index,val) ;
}

