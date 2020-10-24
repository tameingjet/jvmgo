package classfile

import "math"
type ConstantIntegerInfo struct {
	val int32
	//正好容纳一个int常量
	//实际上比int更小的boolean，byte，short，char类型的常量也放在这
}
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
	//4字节
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader){
//读取uint32，转换成float32
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	//先读取一个uint64把它转型成int64类型
	bytes:=reader.readUint64()
	self.val = int64(bytes)
}
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
