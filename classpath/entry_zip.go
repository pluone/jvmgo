package classpath

import "path/filepath"
import "archive/zip"
import "io/ioutil"
import "errors"

//ZipEntry jar或者zip类型
type ZipEntry struct {
	zipFilePath string
}

func newZipEntry(zipFilePath string) *ZipEntry {
	zipFilePath, err := filepath.Abs(zipFilePath)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{zipFilePath}
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(zipEntry.zipFilePath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, zipEntry, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (zipEntry *ZipEntry) String() string {
	return zipEntry.zipFilePath
}
