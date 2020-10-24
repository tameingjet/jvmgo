package rtda
//帧
type Frame struct {
	lower *Frame ;//实现链表数据结构
	LocalVars LocalVars ;//局部变量表指针
	OperandStack *OperandStack ;//操作数栈指针
	Thread  *Thread ;
 }

func NewFrame(maxLocals,maxStack uint) *Frame{
	//创造Frame实例
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStack:  newOperandStack(maxStack),
	}
}



