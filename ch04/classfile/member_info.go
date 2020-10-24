package classfile
type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}
func readMembers(reader *ClassReader,cp ConstantPool) []*MemberInfo{
	//读取字段表和代码表
	memberCount := reader.readUint16()
	members := make([]*MemberInfo,memberCount)
	for i:= range members {
		members[i] = readMember(reader,cp)
	}
	return members
}
func readMember(reader *ClassReader,cp ConstantPool)  *MemberInfo{
	//读取字段和方法数据
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader,cp),
	}
}
func (self *MemberInfo)AccessFlags() uint16{
	return self.accessFlags ;
}//getter
func (self *MemberInfo)Name() string{
	//从常量池查找字段或方法名
	return self.cp.getUtf8(self.nameIndex)
}//getter
func (self *MemberInfo)Descriptor() string{
	//从常量池查找字段或方法描述池
	return self.cp.getUtf8(self.descriptorIndex)
}//getter
