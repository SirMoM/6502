// Code generated by "stringer -type=Instruction -trimprefix=_"; DO NOT EDIT.

package computer

import (
	"fmt"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LDA_I-169]
	_ = x[LDX_I-162]
	_ = x[ADC_ZX-117]
	_ = x[JMP_ABS-76]
	_ = x[JMP_IND-108]
}

const (
	_Instruction_name_0 = "JMP_ABS"
	_Instruction_name_1 = "JMP_IND"
	_Instruction_name_2 = "ADC_ZX"
	_Instruction_name_3 = "LDX_I"
	_Instruction_name_4 = "LDA_I"
)

func (i Instruction) String() string {
	switch {
	case i == 76:
		return _Instruction_name_0
	case i == 108:
		return _Instruction_name_1
	case i == 117:
		return _Instruction_name_2
	case i == 162:
		return _Instruction_name_3
	case i == 169:
		return _Instruction_name_4
	default:
		return "Instruction(" + fmt.Sprintf("%#02X", uint8(i)) + ")"
	}
}
