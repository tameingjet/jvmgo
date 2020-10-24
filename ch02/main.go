package main

import (
	"flag"
	"fmt"
	"jvmgo/ch02/classpath"
	"os"
	"strings"
	)
func startJVM(cmd *Cmd)  {
	//fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
	cp:=classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",cp,cmd.class,cmd.args)
	className := strings.Replace(cmd.class,".","/",-1)
	classData,_,err := cp.ReadClass(className)//读取主类数据打印到控制台
	if err!=nil {
		fmt.Printf("Could not found or load main class %s\n",cmd.class)
		return
	}
	fmt.Printf("class data:%v\n",classData)

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
