package loads
import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//加载指令从局部变量表获取变量，然后推入操作数栈顶
//aload  操作引用类型
//xaload 操作数组
//FLOAD 操作int变量
type FLOAD struct {
	base.Index8Instruction ;
}
type FLOAD_0 struct {
	base.NoOperandsInstruction ;
}
type FLOAD_1 struct {
	base.NoOperandsInstruction ;
}
type FLOAD_2 struct {
	base.NoOperandsInstruction ;
}
type FLOAD_3 struct {
	base.NoOperandsInstruction ;
}
//避免重复代码，定义一个函数供FLOAD使用
func _fload(frame *rtda.Frame,index uint)  {
	val := frame.LocalVars.GetFloat(index);
	frame.OperandStack.PushFloat(val) ;
}
//其余4条指令的索引隐含在操作码中
func (self *FLOAD) Execute(frame *rtda.Frame)  {
	_fload(frame,uint(self.Index));
}
func (self *FLOAD_0) Execute(frame *rtda.Frame){
	_fload(frame,0);
}
func (self *FLOAD_1) Execute(frame *rtda.Frame){
	_fload(frame,1);
}
func (self *FLOAD_2) Execute(frame *rtda.Frame){
	_fload(frame,2);
}
func (self *FLOAD_3) Execute(frame *rtda.Frame){
	_fload(frame,3);
}
