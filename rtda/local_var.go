package rtda

import "math"
import "jvmgo/rtda/heap"

//LocalVar 局部变量表类型
type LocalVar []heap.Slot

func newLocalVal(maxSize uint) LocalVar {
	if maxSize > 0 {
		return make([]heap.Slot, maxSize)
	}
	return LocalVar{}
}

func (lv LocalVar) SetInt(index uint, val int32) {
	lv[index].SetNum(val)
}

func (lv LocalVar) GetInt(index uint) int32 {
	return lv[index].Num()
}

func (lv LocalVar) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	lv[index].SetNum(int32(bits))
}

func (lv LocalVar) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(lv[index].Num()))
}

func (lv LocalVar) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.set64bitVal(bits, index)
}

func (lv LocalVar) GetDouble(index uint) float64 {
	low := uint32(lv[index].Num())
	high := uint32(lv[index+1].Num())
	return math.Float64frombits(uint64(low) | uint64(high)<<32)
}

func (lv LocalVar) set64bitVal(val uint64, index uint) {
	low := uint32(val)
	high := uint32(val >> 32)
	lv[index].SetNum(int32(low))
	lv[index+1].SetNum(int32(high))
}

func (lv LocalVar) SetLong(index uint, val int64) {
	lv.set64bitVal(uint64(val), index)
}

func (lv LocalVar) GetLong(index uint) int64 {
	low := uint32(lv[index].Num())
	high := uint32(lv[index+1].Num())
	return int64(uint64(low) | uint64(high)<<32)
}

func (lv LocalVar) SetRef(index uint, ref *heap.Object) {
	lv[index].SetRef(ref)
}

func (lv LocalVar) GetRef(index uint) *heap.Object {
	return lv[index].Ref()
}
