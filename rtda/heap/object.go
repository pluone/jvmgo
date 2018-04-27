package heap

type Object struct{
	class *Class
	instanceFields Slots
}

func (object *Object) Class() *Class{
	return object.class
}

func (object *Object) InstanceFields() Slots{
	return object.instanceFields
}
func (object *Object) IsInstanceOf(class *Class) bool{
	return class.isAssignableFrom(object.class)
}

