package rtda

//Slot 操作数栈和局部变量表中存储的基本单元
type Slot struct{
	num int32
	ref *Object
}

type Object struct {
	// todo
}