package rtda
//局部变量表按索引访问
//实现局部变量表
type Slot struct {
	num int32 ; //存放整数
	ref *Object ;//存放引用
}

