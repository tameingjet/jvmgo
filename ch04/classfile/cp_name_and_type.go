package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex uint16 // 字段和方法名
	descriptorIndex uint16 // 字段或方法描述
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	//索引指向utf-8常量
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
