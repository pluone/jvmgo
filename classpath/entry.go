package classpath

import "strings"
import "os"

const pathListSeperator = string(os.PathListSeparator)

//Entry classpath组成类型的抽象
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
