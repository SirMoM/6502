package programs

import c "noah-ruben.com/6502/computer"

var MiniProg MiniProgram

func init() {
	MiniProg = MiniProgram{data: []c.Word{
		c.Word(c.LDA_I),
		0xF9,
		c.Word(c.ADC_ZX),
		0x0F,
	}}
}

type MiniProgram struct {
	data []c.Word
}

// CopyToMemory copies the program to memory:
// 1. Loads program start address into the start vector of the 6502
// 2. Loads program data into memory
func (m MiniProgram) CopyToMemory(addr c.Address, mem c.Memory) error {

	// Load program start address into the start vector of the 6502
	mem.WriteAddress(c.Address(0xFFFC), addr)

	// Load program data into memory
	for idx, word := range m.data {
		addrWithOffset := addr + c.Address(idx)
		mem.WriteWord(addrWithOffset, word)
	}

	return nil
}
