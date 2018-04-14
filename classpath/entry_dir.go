package classpath

import "path/filepath"
import "io/ioutil"

//DirEntry 文件夹类型
type DirEntry struct {
	absPathDirName string
}

func newDirEntry(path string) *DirEntry {
	absPathDirName, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absPathDirName}
}

func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	classFileName := filepath.Join(dirEntry.absPathDirName, className)
	data, err := ioutil.ReadFile(classFileName)
	return data, dirEntry, err
}

func (dirEntry *DirEntry) String() string {
	return dirEntry.absPathDirName
}
