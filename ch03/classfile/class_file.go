package classfile

//魔数-版本号-常量池-类访问标志-常量池索引-接口索引表-字段表，方法表
type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
    constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fileds []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}
func Parse(classData []byte)(cf *ClassFile,err error){
	//把[]byte解析为ClassFile结构体
	//defer func() {
	//	if r := recover();r!=nil{
	//		var ok bool
	//		err,ok = r.(error)
	//		if !ok {
	//			fmt.Println("!!!!!!!!!")
	//			err = fmt.Errorf("%v",r)
	//			 }
	//	}
	//}()
	cr:=&ClassReader{classData}
	cf =&ClassFile{}
	cf.read(cr)
	return
}
func (self *ClassFile)read(reader *ClassReader)  {
	//调用其他方法解析Class文件
	self.readandCheckMagic(reader)
	self.readandCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass =reader.readUint16()
	self.superClass =reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fileds = readMembers(reader,self.constantPool)
	self.methods = readMembers(reader,self.constantPool)
	self.attributes = readAttributes(reader,self.constantPool)
}
func(self *ClassFile) readandCheckMagic(reader *ClassReader){
	//魔数：每个不同的文件格式会以固定的几个字节开头，叫做魔数，如pdf的是0x25,0x50
	//class文件的魔数：0xCAFEBABE
	magic:=reader.readUint32()
	if magic!=0xCAFEBABE {
		//class文件不符合格式
		panic("java.lang.ClassFormatError:magic!")
	}
}
func(self *ClassFile) readandCheckVersion(reader *ClassReader){
	//魔数后是class文件的此版本号和主版本号（M.m）
	//oracle实现虚拟机完全向后兼容版本
	self.minorVersion=reader.readUint16()
	self.majorVersion=reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0{
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}
func(self *ClassFile)MinorVersion() uint16{
	return self.minorVersion
}//getter
func(self *ClassFile)MajorVersion() uint16{
	return self.majorVersion
}//getter
func(self *ClassFile) ConstantPool()ConstantPool{
	return self.constantPool
}//getter
func(self *ClassFile) AccessFlags() uint16{
	return self.accessFlags
}//getter
func(self *ClassFile) Fileds() []*MemberInfo{
	return self.fileds
}//getter
func(self *ClassFile) Methods() []*MemberInfo{
	return self.methods
}//getter
func(self *ClassFile) ClassName() string{
	//常量池查找类名
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName()string{
	//常量池查找超类名
	if self.superClass>0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "java.lang.object"///???? 73页
}
func (self *ClassFile) InterfaceName() []string{
	interfaceNames:= make([]string,len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}