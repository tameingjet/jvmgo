package classpath

import (
	"os"
	"strings"
)
//实现类路径
//分为启动类路径，拓展类路径，用户类路径
const pathListSeparator  = string(os.PathListSeparator)//存放路径分割符

type Entry interface {
	readClass(className string) ([]byte,Entry,error)//寻找和加载class的方法
	//如输入java/lang，就是读取java.lang
	String() string //tostring
}

func newEntry(path string) Entry  {
	//根据参数不同创造不同的类型的Entry实例
	if strings.Contains(path,pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*"){
	//检测后缀（文件名.*）
	return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		//检测后缀（文件名.jar  文件名.zip）
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
