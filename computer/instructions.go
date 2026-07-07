//go:generate stringer -type=Instruction

package computer

// Instruction represents a 6502 CPU instruction
type Instruction uint8

//goland:noinspection ALL
const (
	// Load Accumulation
	// LDA_I Load Immediate
	LDA_I Instruction = 0xA9
	// LDA_Z Load thru Zero Page
	LDA_Z Instruction = 0xA5

	// Load X-Register
	LDX_I Instruction = 0xA2

	// Store Accumulator
	STA_Z  Instruction = 0x85 // Zero Page
	STA_ZX Instruction = 0x95 // Zero Page X
	STA_A  Instruction = 0x8D // Absolute
	STA_AX Instruction = 0x9D // Absolute X
	STA_AY Instruction = 0x99 // Absolute Y
	STA_IX Instruction = 0x9D // Indirect X
	STA_IY Instruction = 0x99 // Indirect Y

	// ADC - Add with Carry
	ADC_Z  Instruction = 0x65 // Zero Page
	ADC_ZX Instruction = 0x75 // Zero Page,X

	// Shift and status
	ASL_A Instruction = 0x0A // Accumulator
	CLC   Instruction = 0x18 // Clear Carry

	// Return
	RTS Instruction = 0x60

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
