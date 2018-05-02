package instructions

import (
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions/comparisons"
	"jvmgo/instructions/constants"
	"jvmgo/instructions/control"
	"jvmgo/instructions/loads"
	"jvmgo/instructions/math"
	"jvmgo/instructions/stack"
	"jvmgo/instructions/stores"
	"jvmgo/instructions/references"
)

// NoOperandsInstruction singletons
var (
	nop         = &constants.NOP{}
	aconst_null = &constants.ACONST_NULL{}
	iconst_m1   = &constants.ICONST_M1{}
	iconst_0    = &constants.ICONST_0{}
	iconst_1    = &constants.ICONST_1{}
	iconst_2    = &constants.ICONST_2{}
	iconst_3    = &constants.ICONST_3{}
	iconst_4    = &constants.ICONST_4{}
	iconst_5    = &constants.ICONST_5{}
	lconst_0    = &constants.LCONST_0{}
	lconst_1    = &constants.LCONST_1{}
	fconst_0    = &constants.FCONST_0{}
	fconst_1    = &constants.FCONST_1{}
	fconst_2    = &constants.FCONST_2{}
	dconst_0    = &constants.DCONST_0{}
	dconst_1    = &constants.DCONST_1{}
	iload_0     = &loads.ILOAD_0{}
	iload_1     = &loads.ILOAD_1{}
	iload_2     = &loads.ILOAD_2{}
	iload_3     = &loads.ILOAD_3{}

	aload_0 = &loads.ALOAD_0{}
	aload_1 = &loads.ALOAD_1{}
	aload_2 = &loads.ALOAD_2{}
	aload_3 = &loads.ALOAD_3{}

	istore_1 = &stores.ISTORE_1{}
	istore_2 = &stores.ISTORE_2{}
	astore_0 = &stores.ASTORE_0{}
	astore_1 = &stores.ASTORE_1{}
	astore_2 = &stores.ASTORE_2{}
	astore_3 = &stores.ASTORE_3{}
)

func NewInstruction(opcode uint8) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x10:
		return &constants.BIPUSH{}
	case 0x12:
		return &constants.LDC{}
	case 0x13:
		return &constants.LDC_W{}
	case 0x14:
		return &constants.LDC2_W{}
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x59:
		return &stack.DUP{}
	case 0x60:
		return &math.IADD{}
	case 0x84:
		return &math.IINC{}
	case 0x99:
		return &comparisons.IFEQ{}
	case 0x9a:
		return &comparisons.IFNE{}
	case 0x9b:
		return &comparisons.IFLT{}
	case 0x9c:
		return &comparisons.IFGE{}
	case 0x9d:
		return &comparisons.IFGT{}
	case 0x9e:
		return &comparisons.IFLE{}
	case 0xa3:
		return &comparisons.IF_ICMPGT{}
	case 0xa7:
		return &control.GOTO{}
	case 0xb2:
		return &references.GetStatic{}
	case 0xb3:
		return &references.PutStatic{}
	case 0xb4:
		return &references.GetField{}
	case 0xb5:
		return &references.PutField{}
	case 0xb6:
		return &references.INVOKE_VIRTUAL{}
	case 0xb7:
		return &references.InvokeSpecial{}
	case 0xbb:
		return &references.NEW{}
	case 0xc0:
		return &references.CheckCast{}
	case 0xc1:
		return &references.InstanceOf{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
