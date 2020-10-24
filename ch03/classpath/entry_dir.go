package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//目录形式的类路径
type DirEntry struct {
	absDir string//存放目录的绝对路径
}
func newDirEntry(path string) *DirEntry{
	//构造direntry,把参数转换成绝对路径
	absDir,err := filepath.Abs(path)
	if err!= nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func(self *DirEntry) readClass(className string) ([]byte,Entry,error){
	//方法匹配即可实现接口
	fileName := filepath.Join(self.absDir,className)//拼成完整文件名
	data,err:=ioutil.ReadFile(fileName)//读取class内容
	return data,self,err
}

func (self *DirEntry) String() string{
	return  self.absDir
}