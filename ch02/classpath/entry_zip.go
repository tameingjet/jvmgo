package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}
func newZipEntry(path string) *ZipEntry{
	absDir,err := filepath.Abs(path)
	if err!= nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}
func(self *ZipEntry) readClass(className string) ([]byte,Entry,error){
    //zip文件中提取class文件
    r,err := zip.OpenReader(self.absPath)//打开zip文件
    if err != nil{
    	return nil,nil,err
	}
	defer r.Close()
	for _, f := range r.File {//遍历找class文件
		if f.Name == className {
			rc,err := f.Open()
			if err != nil{
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)//找到读出来
			if err != nil{
				return nil,nil,err
			}
			return data,self,nil
		}
	}
	return nil,nil,errors.New("class not found: "+className)
}

func (self *ZipEntry) String() string{
	return self.absPath
}



