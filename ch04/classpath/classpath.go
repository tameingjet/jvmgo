package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	//存放三种路径
	bootClasspath Entry//启动类路径
	extClasspath Entry//扩展路径
	userClasspath Entry//用户路径
}
func Parse(jreOption,cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)//使用-Xjre解析启动类路径和扩展类路径
	cp.parseUserClasspath(cpOption)
	return cp
}
func (self *Classpath)parseBootAndExtClasspath(jreOption string)  {
	jreDir := getJreDir(jreOption)//找jre目录
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//jre/lib/next
	jreExtPath:=filepath.Join(jreDir,"lib","ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}
func (self *Classpath)parseUserClasspath(cpOption string) {
	if cpOption == ""{
		cpOption = "."
	}//如果用户没有提供-Classpath/-cp选项，则使用当前目录作为用户类路径
	self.userClasspath = newEntry(cpOption)
}
func getJreDir(jreOption string)string{
	//优先使用用户输入的-Xjre作为jre目录，如果没有在当前目录下找jre，还找不到
	//使用JAVA_HOME环境变量
	if jreOption != ""&&exists(jreOption){
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME");jh!=""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}
func exists(path string) bool{
	//判断目录是否存在
	if _,err := os.Stat(path);err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath)ReadClass(className string)([]byte,Entry,error){
	//搜索class文件
	className = className+".class"
	if data,entry,err:=self.bootClasspath.readClass(className);err == nil {
		return data,entry,err
	}
	if data,entry,err := self.extClasspath.readClass(className);err == nil {
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}
func(self *Classpath) String() string{
	return self.userClasspath.String()
}