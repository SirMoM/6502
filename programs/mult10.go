package programs

import c "noah-ruben.com/6502/computer"

var Mult10Prog Mult10Program

func init() {
	Mult10Prog = Mult10Program{data: []any{
		c.ASL_A,
		c.STA_Z,
		0x10,
		c.ASL_A,
		c.ASL_A,
		c.CLC,
		c.ADC_Z,
		0x10,
		c.RTS,
	}}
}

type Mult10Program struct {
	data []any
}

// CopyToMemory copies the multiply-by-10 routine to memory and seeds the temp byte.
func (m Mult10Program) CopyToMemory(addr c.Address, mem c.Memory) error {
	mem.WriteAddress(c.Address(0xFFFC), addr)

	for idx, word := range m.data {
		mem.WriteWord(addr+c.Address(idx), c.ToWord(word))
	}

	mem.WriteWord(c.Address(0x0010), c.Word(0x00))

	return nil
}
