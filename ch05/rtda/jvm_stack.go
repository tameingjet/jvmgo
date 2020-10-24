package rtda

type Stack struct {
	maxSize uint ;//最多可以容纳多少帧
	size uint;//当前栈的大小
	_top *Frame ;//保存栈顶指针
}

func newStack(maxSize uint) *Stack  {
	return &Stack{
		maxSize: maxSize,
	}
}
func (self *Stack) push(frame *Frame){
	//把帧推入栈
	if self.size >= self.maxSize{
		panic("java.lang.StackoverFlowError") ;
	}
	if self._top!=nil {
		frame.lower = self._top ;
	}
	self._top = frame;
	self.size++ ;
}
func (self *Stack) pop() *Frame{
 //把栈顶帧弹出
	if self._top == nil {
		panic("jvm stack is empty") ;
	}
	top := self._top ;
	self._top = top.lower
	top.lower = nil ;
	self.size -- ;
	return top ;
}
func (self *Stack) top() *Frame {
	//返回栈顶帧，但不弹出
	if self._top == nil  {
		panic("jvm stack is empty") ;
	}
	return self._top
}