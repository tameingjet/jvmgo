package classfile

type LineNumberTableEntry struct {
	startPc uint16 ;
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	//存放方法的行号信息
	LineNumberTable []*LineNumberTableEntry
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader){
	lineNumberTableLength := reader.readUint16();
	self.LineNumberTable = make([]*LineNumberTableEntry,lineNumberTableLength);
	for i := range self.LineNumberTable {
		self.LineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}