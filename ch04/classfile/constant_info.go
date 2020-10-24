package classfile
//每种常量的格式不同，常量的第一字节（tag）用来区分类型
const(
	CONSTANT_Class                =7
	CONSTANT_Fieldref             =9
	CONSTANT_Methodref            =10
	CONSTANT_InterfaceMethodref   =11
	CONSTANT_String               =8
	CONSTANT_Integer              =3
	CONSTANT_Float                =4
	CONSTANT_Long                 =5
	CONSTANT_Double               =6
	CONSTANT_NameAndType          =12
	CONSTANT_Utf8                 =1
	CONSTANT_MethodHandle         =15
	CONSTANT_Methodtype           =16
	CONSTANT_InvokeDynamic        =18
	)
