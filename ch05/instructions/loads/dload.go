package loads
import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//加载指令从局部变量表获取变量，然后推入操作数栈顶
//aload  操作引用类型
//xaload 操作数组
//iload 操作int变量
type DLOAD struct {
	base.Index8Instruction ;
}
type DLOAD_0 struct {
	base.NoOperandsInstruction ;
}
type DLOAD_1 struct {
	base.NoOperandsInstruction ;
}
type DLOAD_2 struct {
	base.NoOperandsInstruction ;
}
type DLOAD_3 struct {
	base.NoOperandsInstruction ;
}
//避免重复代码，定义一个函数供iload使用
func _dload(frame *rtda.Frame,index uint)  {
	val := frame.LocalVars.GetDouble(index);
	frame.OperandStack.PushDouble(val) ;
}
//其余4条指令的索引隐含在操作码中
func (self *DLOAD) Execute(frame *rtda.Frame)  {
	_dload(frame,uint(self.Index));
}
func (self *DLOAD_0) Execute(frame *rtda.Frame){
	_dload(frame,0);
}
func (self *DLOAD_1) Execute(frame *rtda.Frame){
	_dload(frame,1);
}
func (self *DLOAD_2) Execute(frame *rtda.Frame){
	_dload(frame,2);
}
func (self *DLOAD_3) Execute(frame *rtda.Frame){
	_dload(frame,3);
}
