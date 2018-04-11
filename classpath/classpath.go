package classpath
import "path/filepath"
import "os"
import "errors"
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption,cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption == ""{
		cpOption = "."
	}
	self.userClasspath = newCompositeEntry(cpOption)
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)
	self.bootClasspath = newWildcardEntry( filepath.Join(jreDir,"lib","*"))
	self.extClasspath = newWildcardEntry(filepath.Join("lib","ext","*"))
}

func getJreDir(jreOption string) string{
	if jreOption!="" && exist(jreOption){
		return jreOption
	}
	if exist(filepath.Join(".","jre")){
		return filepath.Join(".","jre")
	}
	if jh:= os.Getenv("JAVA_HOME") ; jh!=""{
		return filepath.Join(jh,"jre")
	}
	panic("cannot find jre folder!")

}

func exist(path string) bool {
	_,err:=os.Stat(path)
	if err!=nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte,Entry,error){
	className = className + ".class"
	data,entry,err := self.bootClasspath.readClass(className)
	if err == nil {
		return data,entry,err
	}
	data,entry,err = self.extClasspath.readClass(className)
	if err == nil {
		return data,entry,err
	}
	data,entry,err = self.userClasspath.readClass(className)
	if err == nil {
		return data,entry,err
	}
	return nil,nil,errors.New("class not found, className: "+className)
}

func (self *Classpath) String() string {
	return self.bootClasspath.String()+"\n"+self.extClasspath.String()+"\n"+self.userClasspath.String()
}