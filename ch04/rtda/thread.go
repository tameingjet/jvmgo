package rtda

type Thread struct {
	pc int ;//每个线程都有自己的pc寄存器
	stack *Stack ;//Java虚拟机栈
	//Java虚拟机可以是连续的空间可以不连续
	//可以是固定大小，也可以在运行时动态扩展
//	如果有限制大小但执行线程所需空间超过限制，会抛出StackOverFlow异常
//如果可以动态扩展，但内存耗尽，会导致OutOfMemoryError异常
}

func NewThread() *Thread  {
	return &Thread{stack:newStack(1024)} ;
	//创建Stack结构体实例，参数表示创建Stack最多可以容纳多少帧
}

func (self *Thread) PC() int{
	return self.pc
}
func (self *Thread) SetPC(pc int)  {
	self.pc = pc ;
}
func (self *Thread) PushFrame(frame *Frame)  {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame{
	//返回当前帧
	return self.stack.top()
}


