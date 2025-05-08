//go:generate stringer -type=Instruction

package computer

// Instruction represents a 6502 CPU instruction
type Instruction uint8

//goland:noinspection ALL
const (
	// Load Accumulation
	LDA_I Instruction = 0xA9
	LDX_I Instruction = 0xA2

	// ADC - Add with Carry
	ADC_ZX Instruction = 0x75 // Zero Page,X

	// JMP - Jump
	JMP_ABS Instruction = 0x4C // Absolute
	JMP_IND Instruction = 0x6C // Indirect
)

// Bit masks for processor status flags
const (
	bit7 Word = 0x80
	bit6 Word = 0x40
	bit5 Word = 0x20
	bit4 Word = 0x10
	bit3 Word = 0x8
	bit2 Word = 0x4
	bit1 Word = 0x2
	bit0 Word = 0x1
)
