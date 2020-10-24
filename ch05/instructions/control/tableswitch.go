package control
//Java语言中的switch-case语句有两种实现方式：如果case值可以
//编码成一个索引表，则实现成tableswitch指令；否则实现成
//lookupswitch指令。
import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
//defaultOffset对应默认情况下执行跳转所需的字节码偏移量
//low和high记录case的取值范围；jumpOffsets是一个索引表，里面存
//放high-low+1个int值，对应各种case情况下，执行跳转所需的字节
//码偏移量

type TABLE_SWITCH struct {
	defaultOffset int32 ;
	low int32 ;
	high int32;
	jumpOffsets []int32 ;
}
//tableSwitch指令操作码后面有0-3字节的padding(SkippPadding())
//以保证defaultOffset在字节码中的地址是4的倍数
func (self *TABLE_SWITCH)FetchOperands(reader *base.BytecodeReader)  {
	reader.SkipPadding() ;
	self.defaultOffset = reader.ReadInt32() ;
	self.low = reader.ReadInt32() ;
	self.high = reader.ReadInt32() ;
	jumpOffsetsCount := self.high - self.low+1 ;
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount);
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame)  {
	index := frame.OperandStack.PopInt()
	var offset int ;
	if index>=self.low&&index<=self.high {
		offset = int(self.jumpOffsets[index]-self.low)
		} else {
			offset = int(self.defaultOffset) ;
	}
	base.Branch(frame,offset) ;
}