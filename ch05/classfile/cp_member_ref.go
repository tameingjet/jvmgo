package classfile

type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16//常量池索引,指向class常量
	nameAndTypeIndex uint16//指向NameAndType常量
}
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader){
	self.classIndex = reader.readUint16()//
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo)ClassName() string  {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor()(string,string)  {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
//go语言没有继承，可以通过嵌套结构体实现下面三个常量
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}