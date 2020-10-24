package classfile

type ConstantPool []ConstantInfo
//常量池：1.表头给的常量池大小比实际大1
//		2.有效的常量池索引是1~n-1，0是无效索引
//		3.CONSTANT_Long_Info,CONSTANT_Double_info各占两个位置
//分为两类：1.字面量:数字常量和字符串常量
//		  2.符号：类和接口名，字段和方法信息等--通过索引指向utf-8常量
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
func readConstantPool(reader *ClassReader) ConstantPool{
	//读取常量池
	cpCount:=int(reader.readUint16())
	cp:=make([]ConstantInfo,cpCount)
	for i := 1; i < cpCount; i++ {//索引从1开始
		cp[i]=readConstantInfo(reader,cp)
		switch cp[i].(type) {
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++//占两个位置
			}
	}
	return cp
}
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	//按索引查找常量
	if cpInfo := self[index];cpInfo!=nil{
		return cpInfo;
	}
	panic("Invalid constant pool index!")
}
func (self ConstantPool) getNameAndType(index uint16) (string,string){
	//从常量池查找字段或方法的名字和描述符
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name:=self.getUtf8(ntInfo.nameIndex)
	_type:=self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}
func (self ConstantPool) getClassName(index uint16) string{
	//从常量池查找类名
	classInfo:=self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
func (self ConstantPool) getUtf8(index uint16) string{
	//从常量池查找utf-8字符串
	utf8Info:=self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
func readConstantInfo(reader *ClassReader,cp ConstantPool) ConstantInfo{
	//先读出tag值，调用newConstantInfo()创建具体的常量，调用readInfo()方法读取常量信息
	tag:=reader.readUint8()
	c:=newConstantInfo(tag,cp)
	c.readInfo(reader)
	return c
}
func newConstantInfo(tag uint8,cp ConstantPool) ConstantInfo{
	switch tag {
	case CONSTANT_Class:return &ConstantClassInfo{cp:cp}
	case CONSTANT_Fieldref:return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp:cp}}
	case CONSTANT_Methodref:return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp:cp}}
	case CONSTANT_InterfaceMethodref:return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp:cp}}
	case  CONSTANT_String:return &ConstantStringInfo{cp:cp}
	case CONSTANT_Integer:return &ConstantIntegerInfo{}
	case CONSTANT_Float:return &ConstantFloatInfo{}
	case CONSTANT_Long:return &ConstantLongInfo{}
	case CONSTANT_Double:return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:return &ConstantMethodHandleInfo{}//看源码
	case CONSTANT_Methodtype:return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError:constant pool tag!")
	}
}


