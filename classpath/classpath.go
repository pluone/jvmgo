package classpath

import "path/filepath"
import "os"
import "errors"

//Classpath classpath
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//Parse do parse
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClasspath = newCompositeEntry(cpOption)
}

func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	classpath.bootClasspath = newWildcardEntry(filepath.Join(jreDir, "lib", "*"))
	classpath.extClasspath = newWildcardEntry(filepath.Join("lib", "ext", "*"))
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exist(jreOption) {
		return jreOption
	}
	if exist(filepath.Join(".", "jre")) {
		return filepath.Join(".", "jre")
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("cannot find jre folder!")

}

func exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//ReadClass read class data
func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	data, entry, err := classpath.bootClasspath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	data, entry, err = classpath.extClasspath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	data, entry, err = classpath.userClasspath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	return nil, nil, errors.New("class not found, className: " + className)
}

func (classpath *Classpath) String() string {
	return classpath.userClasspath.String()
}
