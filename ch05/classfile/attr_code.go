package classfile

type CodeAttribute struct {
	//存放字节码
	cp ConstantPool ;
	maxStack uint16 ; //给出操作栈的最大深度
	maxLocals uint16 ;//给出局部变量表的大小
	code []byte ;//字节码
	exceptionTable []*ExceptionTableEntry ;//异常处理表
	attributes []AttributeInfo;//属性表
}
type ExceptionTableEntry struct {
	startPc uint16 ;
	endPc uint16 ;
	handlePc uint16 ;
	catchType uint16 ;
}

func (self *CodeAttribute) readInfo(reader *ClassReader)  {
	self.maxStack = reader.readUint16() ;
	self.maxLocals = reader.readUint16();
	codeLength := reader.readUint32() ;
	self.code = reader.readBytes(codeLength);
	self.exceptionTable = readExceptionTable(reader)

}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
     exceptionTableLength := reader.readUint16();
     exceptionTable := make([]*ExceptionTableEntry,exceptionTableLength);
     for i:= range exceptionTable{
     	exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlePc:  reader.readUint16(),
			catchType: reader.readUint16(),
		}
	 }
	 return  exceptionTable ;
}