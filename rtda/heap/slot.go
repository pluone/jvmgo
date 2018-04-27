package heap

//Slot 操作数栈和局部变量表中存储的基本单元
type Slot struct {
	num int32
	ref *Object
}

func (slot *Slot) Num() int32 {
	return slot.num
}

func (slot *Slot) Ref() *Object {
	return slot.ref
}

func (slot *Slot) SetNum(val int32) {
	slot.num = val
}

func (slot *Slot) SetRef(ref *Object) {
	slot.ref = ref
}
