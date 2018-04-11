package classpath

import "path/filepath"
import "io/ioutil"

type DirEntry struct{
	absPathDirName string
}

func newDirEntry(path string) *DirEntry{
	absPathDirName,err :=  filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	return &DirEntry{absPathDirName}
}

func (self *DirEntry) readClass(className string)  ([]byte,Entry,error){
	classFileName := filepath.Join(self.absPathDirName,className)
	data,err := ioutil.ReadFile(classFileName)
	return data,self,err
}

func (self *DirEntry) String() string{
	return self.absPathDirName
}