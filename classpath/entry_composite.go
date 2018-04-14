package classpath

import "strings"
import "errors"

//CompositeEntry 复合类型
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeperator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (compositeEntry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range compositeEntry {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found, className: " + className)
}

func (compositeEntry CompositeEntry) String() string {
	strArray := make([]string, len(compositeEntry))
	for _, entry := range compositeEntry {
		strArray = append(strArray, entry.String())
	}
	return strings.Join(strArray, pathListSeperator)
}
