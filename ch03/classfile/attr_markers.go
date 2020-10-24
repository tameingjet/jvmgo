package classfile
//
type MarkerAttribute struct {

}
type DeprecatedAttribute struct {
	//@Deprecated用于指出类，接口，方法不建议使用
	MarkerAttribute ;
}
type SyntheticAttribute struct {
	//用于支持嵌套类和嵌套接口
	MarkerAttribute ;
}

func (self *MarkerAttribute)readInfo(reader *ClassReader)  {
	//readNoting
}