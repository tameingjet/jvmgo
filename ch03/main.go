package main

import (
	"flag"
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
	"os"
	"strings"
)
func startJVM(cmd *Cmd)  {
	//fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
	cp:=classpath.Parse(cmd.XjreOption,cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
	cf :=loadClass(className,cp);
	fmt.Println(cmd.class);
	printClassInfo(cf);

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
