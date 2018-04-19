package rtda

import "math"

//OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStackSize uint) *OperandStack {
	if maxStackSize > 0 {
		return &OperandStack{
			size:  0,
			slots: make([]Slot, maxStackSize),
		}
	}
	return nil
}

//PushInt 将int32类型入栈
func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.size++
}

//PopInt int32类型出栈
func (stack *OperandStack) PopInt() int32 {
	stack.size--
	val := stack.slots[stack.size].num
	return val
}

//PushFloat float32类型入栈
func (stack *OperandStack) PushFloat(val float32) {
	stack.slots[stack.size].num = int32(math.Float32bits(val))
	stack.size++
}

//PopFloat float32类型出栈
func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	bytes := stack.slots[stack.size].num
	return math.Float32frombits(uint32(bytes))
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].num = int32(val)
	stack.slots[stack.size+1].num = int32(val >> 32)
	stack.size += 2
}

func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	low := uint32(stack.slots[stack.size].num)
	high := uint32(stack.slots[stack.size+1].num)
	return int64(low) | int64(high)<<32
}

func (stack *OperandStack) PushDouble(val float64) {
	stack.PushLong(int64(math.Float64bits(val)))
}

func (stack *OperandStack) PopDouble() float64 {
	bytes := uint64(stack.PopLong())
	return math.Float64frombits(bytes)
}

func (stack *OperandStack) PushRef(val *Object) {
	stack.slots[stack.size].ref = val
	stack.size++
}

func (stack *OperandStack) PopRef() *Object {
	stack.size--
	refVal := stack.slots[stack.size].ref
	stack.slots[stack.size].ref = nil //help gc
	return refVal
}
