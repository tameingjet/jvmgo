package main

import (
	"flag"
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
	"jvmgo/ch04/rtda"
	"os"
)
func startJVM(cmd *Cmd)  {
	//fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
	frame := rtda.NewFrame(100,100);
	testLocalVars(frame.LocalVars);
	testOperandStack(frame.OperandStack);

}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.Popref())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0,100);
	vars.SetInt(1,-100);
	vars.SetLong(2,2997924580);
	vars.SetLong(4,-2997924580);
	vars.SetFloat(6,3.1415926);
	vars.SetDouble(7,2.71828182845);
	vars.SetRef(9,nil);
	fmt.Println(vars.GetInt(0));
	fmt.Println(vars.GetInt(1));
	fmt.Println(vars.GetLong(2));
	fmt.Println(vars.GetLong(4));
	fmt.Println(vars.GetFloat(6));
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9));
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:  %v.%v\n",cf.MajorVersion(),cf.MinorVersion());
	fmt.Printf("constant count: %v\n",len(cf.ConstantPool()));
	fmt.Printf("access flags:0x%x\n",cf.AccessFlags());
	fmt.Printf("this class:%v\n",cf.ClassName());
	fmt.Printf("super class : %v\n",cf.SuperClassName());
	fmt.Printf("interfaces:%v\n",cf.InterfaceName());
	fmt.Printf("field count:%v\n",len(cf.Fileds()));
	for _, f := range cf.Fileds() {
		fmt.Printf("   %s\n",f.Name())
	}
}
//读取解析class文件
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
	classData,_,err := cp.ReadClass(className);
	if err!=nil {
		panic(err);
	}
	cf,err:=classfile.Parse(classData);
	if err!=nil {
		panic(err);
	}
	return  cf ;
}
func main() {
	cmd:=parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
		}else if cmd.helpFlag||cmd.class == "" {
		printUsage()
	}else {
		startJVM(cmd)
	}
}
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
	XjreOption string
}

func printUsage()  {
	fmt.Printf("Usage: %s [-options] class [args...]\n",os.Args[0])
}

func parseCmd() *Cmd{
	cmd:=&Cmd{}
	flag.Usage=printUsage
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message")
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")// 加载jre的路径
	flag.Parse()
	args:=flag.Args()
	if len(args)>0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
