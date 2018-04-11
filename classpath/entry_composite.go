package classpath

import "strings"
import "errors"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry{
	compositeEntry := []Entry{}
	
	for _,path := range strings.Split(pathList,pathListSeperator){
		entry := newEntry(path)
		compositeEntry = append(compositeEntry,entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte,Entry,error){
	for _,entry := range self{
		data,from,err := entry.readClass(className)
		if err==nil{
			return data,from,err
		}
	}
	return nil,nil,errors.New("class not found, className: "+className)
}

func (self CompositeEntry) String() string{
	strArray := make([]string,len(self))
	for _,entry := range self {
		strArray = append(strArray,entry.String())
	}
	return strings.Join(strArray,pathListSeperator)
}