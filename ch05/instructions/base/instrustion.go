package base
//指令：操作码+_+操作数 （iconst_0:把常数0推入操作数栈）
//每条指令都以一个单字节的操作码（opcode）开头,
//Java虚拟机最多只能支持256(2的8次方)条指令
//操作码的助记符：类型+操作（iadd：对int进行加法操作）
import "jvmgo/ch04/rtda"
//有很多指令操作数都是类似的，为了避免重复代码，按照操作数类型定义一些结构体
//并实现FetchOperands（）。相当于抽象类。
//接下来具体的指令继承这些结构体，专注实现Execute（）方法
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader);
	//执行指令逻辑
	Execute(frame *rtda.Frame);
}
type NoOperandsInstruction struct {
	//没有操作数的指令
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	//没操作数
}

type BranchInstruction struct {
	//跳转指令：offset存放跳转偏转量
	Offset int;
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16());
}

type Index8Instruction struct {
	//存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
	//把这类指令抽象成该结构体，Index为局部变量索引
	Index uint ;
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index=uint(reader.ReadUint8())
}

type Index16Instruction struct {
	//一些指令需要访问运行时常量池，常量池由两字节操作数给出
	//抽象成该结构体,Index为索引
	Index uint ;
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index=uint(reader.ReadUint16())
}









