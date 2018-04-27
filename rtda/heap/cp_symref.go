package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (symRef *SymRef) ResolvedClass() *Class {
	if symRef.class == nil {
		symRef.resolveClassRef()
	}
	return symRef.class
}

func (symRef *SymRef) resolveClassRef() {
	d := symRef.cp.class
	c := d.loader.LoadClass(symRef.className)
	if !c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError!")
	}
	symRef.class = c
}
