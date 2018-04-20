package rtda

import "math"

//LocalVar 局部变量表类型
type LocalVar []Slot

func newLocalVal(maxSize uint16) LocalVar {
	if maxSize > 0 {
		return make([]Slot, maxSize)
	}
	return LocalVar{}
}

func (lv LocalVar) SetInt(index uint, val int32) {
	lv[index].num = val
}

func (lv LocalVar) GetInt(index uint) int32 {
	return lv[index].num
}

func (lv LocalVar) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	lv[index].num = int32(bits)
}

func (lv LocalVar) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(lv[index].num))
}

func (lv LocalVar) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.set64bitVal(bits, index)
}

func (lv LocalVar) GetDouble(index uint) float64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return math.Float64frombits(uint64(low) | uint64(high)<<32)
}

func (lv LocalVar) set64bitVal(val uint64, index uint) {
	low := uint32(val)
	high := uint32(val >> 32)
	lv[index].num = int32(low)
	lv[index+1].num = int32(high)
}

func (lv LocalVar) SetLong(index uint, val int64) {
	lv.set64bitVal(uint64(val), index)
}

func (lv LocalVar) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return int64(uint64(low) | uint64(high)<<32)
}

func (lv LocalVar) SetRef(index uint, ref *Object) {
	lv[index].ref = ref
}

func (lv LocalVar) GetRef(index uint) *Object {
	return lv[index].ref
}
