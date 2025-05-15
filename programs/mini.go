package programs

import "noah-ruben.com/6502/computer"

var MiniProg MiniProgram

func init() {
	MiniProg = MiniProgram{data: []computer.Word{
		computer.Word(computer.LDA_I),
		0xF9,
		computer.Word(computer.ADC_ZX),
		0x0F,
	}}
}

type MiniProgram struct {
	data []computer.Word
}

func (m MiniProgram) CopyToMemory(addr computer.Address, mem *computer.Memory) error {
	mem.Data[0xFFFD] = 0x00
	mem.Data[0xFFFC] = 0x02

	for idx, word := range m.data {
		addrWithOffset := addr + computer.Address(idx)
		mem.Data[addrWithOffset] = word
	}

	return nil
}
