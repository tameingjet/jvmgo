package base
//code字段存放字节码，pc字段记录读取到哪个字节
type BytecodeReader struct {
	code []byte ;
	pc int ;
}
//避免每次都要创建一个bytecodereader对象，给他定义一个重设方法
func (self *BytecodeReader)Reset(code []byte , pc int) {
	self.code = code ;
	self.pc = pc ;
}
func (self *BytecodeReader) ReadUint8() uint8  {
	i :=self.code[self.pc];
	self.pc++;
	//byte就是uint8
	return i ;
}
func (self *BytecodeReader) ReadInt8() int8  {
	//直接uint8转int8
	return int8(self.ReadUint8());
}
func (self *BytecodeReader) ReadUint16() uint16  {
	//连续读取两个字节
	byte1 := uint16(self.ReadUint8());
	byte2 := uint16(self.ReadUint8());
	return (byte1<<8)|byte2;
}

func (self *BytecodeReader)ReadInt16() int16  {
	//直接转成int16
	return int16(self.ReadUint16());
}

func (self *BytecodeReader) ReadUint32() uint32  {
	//连续读取四个字节
	byte1 := uint32(self.ReadUint8());
	byte2 := uint32(self.ReadUint8());
	byte3 := uint32(self.ReadUint8());
	byte4 := uint32(self.ReadUint8());
	return (byte1<<24)|(byte2<<16)|(byte3<<8)|byte4;
}
func (self *BytecodeReader)ReadInt32() int32 {
	//直接转成int16
	return int32(self.ReadUint32());
}

func (self *BytecodeReader) SkipPadding()  {
	for self.pc!=0 {
		self.ReadUint8();
	}
}
func (self *BytecodeReader) ReadInt32s(n int32) []int32  {
	ints := make([]int32, n);
	for i := range ints{
		ints[i] = self.ReadInt32() ;
	}
	return ints ;
}

