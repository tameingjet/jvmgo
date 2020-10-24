package classfile

import "fmt"

//常量由java虚拟机严格规定的14种，但属性是可以扩展的，不同虚拟机可以定义
//自己的不同属性
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader,cp ConstantPool) []AttributeInfo {
	//读取属性表
	attributesCount:=reader.readUint16();
	fmt.Print(attributesCount)
	fmt.Println("???????????")
	attributes := make([]AttributeInfo,attributesCount);
	for i := range attributes{
		attributes[i] = readAttribute(reader,cp);
	}
	return attributes;

}
func readAttribute(reader *ClassReader,cp ConstantPool) AttributeInfo  {
	//先读取属性名索引，根据它从常量池中找到属性名，然后读取属性长度，调用newAttributeInfo()函数创建具体的属性实例
	attrnameIndex := reader.readUint16();
	fmt.Print(attrnameIndex);
	fmt.Println("!!!!!!!!!!!!!!!!")
	attrName := cp.getUtf8(attrnameIndex);
	attrLen := reader.readUint32();
	attrInfo:=newAttributeInfo(attrName,attrLen,cp);
	attrInfo.readInfo(reader);
	return  attrInfo ;
}
func newAttributeInfo(attrName string,attrLen uint32,cp ConstantPool)AttributeInfo{
	switch attrName {
	case "Code":return &CodeAttribute{cp:cp};
	case "ConstantValue":return &ConstantValueAttribute{};
	case "Deprecated":return &DeprecatedAttribute{};
	case "Exceptions": return &ExceptionsAttribute{};
	case "LineNumberTable": return &LineNumberTableAttribute{};
	//case "LocalVariableTable" : return  &LocalVariableTableAttribute{};
	case "SourceFile": return &SourceFileAttribute{};
	case "Synthetic": return &SyntheticAttribute{} ;
	default:
		return &UnparsedAttribute{attrName,attrLen,nil} ;
	}
}
