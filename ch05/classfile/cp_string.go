package classfile
//CONSTANT_String_info本身并不存放字符串数据,只存放常量池的索引，这个
//索引指向一个CONSTANT_Utf8_info常量
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo)readInfo(reader *ClassReader){
	self.stringIndex=reader.readUint16()
}
func(self *ConstantStringInfo) String()string{
	return self.cp.getUtf8(self.stringIndex)
}