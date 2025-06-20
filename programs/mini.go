package programs

import c "noah-ruben.com/6502/computer"

var MiniProg MiniProgram

func init() {

	// how to fill zero page?
	MiniProg = MiniProgram{data: []c.Word{
		c.Word(c.LDA_Z),
		0xF9,
		c.Word(c.LDX_I),
		0x0F,
		c.Word(c.ADC_ZX),
		0x01,
	}}
}

type MiniProgram struct {
	data []c.Word
}

// CopyToMemory copies the program to memory:
// 1. Loads program start address into the start vector of the 6502
// 2. Loads program data into memory
func (m MiniProgram) CopyToMemory(addr c.Address, mem c.Memory) error {

	// Load the program start address into the start vector of the 6502
	mem.WriteAddress(c.Address(0xFFFC), addr)

	// Load the program data into memory
	for idx, word := range m.data {
		addrWithOffset := addr + c.Address(idx)
		mem.WriteWord(addrWithOffset, word)
	}

	// Data section?
	mem.WriteWord(c.Address(0x10), c.Word(0x09))

	return nil
}
