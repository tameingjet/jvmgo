package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string)CompositeEntry{
	baseDir := path[:len(path)-1]//去掉末尾星号
	compositeEntry := []Entry{}
	walkFn := func(path string,info os.FileInfo,err error) error{
		if err != nil{
			return err
		}
		if info.IsDir()&&path!= baseDir {
			return filepath.SkipDir
		}//跳过子目录
		if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR") {
		   jarEntry := newZipEntry(path)
		   compositeEntry = append(compositeEntry,jarEntry)
		}//根据后缀名选出jar文件
		return nil
	}
	filepath.Walk(baseDir,walkFn)//遍历baseDir，并用walkfn进行处理
	return compositeEntry
}
