package classfile
//与string常量类似，有一个索引指向utf-8常量
type ConstantClassInfo struct {
	cp ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader)  {
	self.nameIndex = reader.readUint16()

}
func (self *ConstantClassInfo) Name() string  {
	return self.cp.getUtf8(self.nameIndex)
}
