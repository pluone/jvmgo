package classpath
import "strings"
import "path/filepath"
import "os"

func newWildcardEntry(wildcardPath string) CompositeEntry{
	baseDir := wildcardPath[:len(wildcardPath)-1]
	compositeEntry := []Entry{}
	walkFunc := func(path string, info os.FileInfo, err error) error{
		if err !=nil {
			return err
		}
		if path != baseDir && info.IsDir(){
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".zip"){
			entry:= newZipEntry(path)
			compositeEntry = append(compositeEntry,entry)
		}
		return nil
	}
	filepath.Walk(baseDir,walkFunc)
	return compositeEntry
}

