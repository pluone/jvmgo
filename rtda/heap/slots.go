package heap

import "math"

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (slots Slots) SetInt(index uint, val int32) {
	slots[index].num = val
}

func (slots Slots) GetInt(index uint) int32 {
	return slots[index].num
}

func (slots Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	slots[index].num = int32(bits)
}

func (slots Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(slots[index].num))
}

func (slots Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	slots.set64bitVal(bits, index)
}

func (slots Slots) GetDouble(index uint) float64 {
	low := uint32(slots[index].num)
	high := uint32(slots[index+1].num)
	return math.Float64frombits(uint64(low) | uint64(high)<<32)
}

func (slots Slots) set64bitVal(val uint64, index uint) {
	low := uint32(val)
	high := uint32(val >> 32)
	slots[index].num = int32(low)
	slots[index+1].num = int32(high)
}

func (slots Slots) SetLong(index uint, val int64) {
	slots.set64bitVal(uint64(val), index)
}

func (slots Slots) GetLong(index uint) int64 {
	low := uint32(slots[index].num)
	high := uint32(slots[index+1].num)
	return int64(uint64(low) | uint64(high)<<32)
}

func (slots Slots) SetRef(index uint, ref *Object) {
	slots[index].ref = ref
}

func (slots Slots) GetRef(index uint) *Object {
	return slots[index].ref
}
