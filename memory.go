package main

import (
	"fmt"
	"math"
	"strings"
)

// Memory represents the memory of the 6502 CPU
type Memory struct {
	Data []Word
}

// Init initializes the memory with default values
func (mem *Memory) Init() {
	mem.Data = make([]Word, math.MaxUint16+1)
	mem.Data[0xFFFC] = 0x00
	mem.Data[0xFFFD] = 0x02

	mem.Data[0x0200] = Word(LDX_I)
	mem.Data[0x0201] = 0xF9
	mem.Data[0x0202] = Word(ADC_ZX)
	mem.Data[0x0203] = 0x0F
}

// String returns a string representation of the memory
func (mem Memory) String() string {
	sb := strings.Builder{}
	blockSize := 32
	collums := 2
	for i := 0; i < math.MaxUint16; i += blockSize {
		fmt.Fprintf(&sb, "── Blocks from %#04X to %#04X ──\n", i, i+blockSize-1)
		for j := 0; j < blockSize/collums; j++ {
			i1 := i + j
			v1 := mem.Data[i1]
			fmt.Fprintf(&sb, "[%#04X] %#02X ─ ", i1, v1)
			i2 := j + i + blockSize/collums
			v2 := mem.Data[i2]
			fmt.Fprintf(&sb, "[%#04X] %#02X\n", i2, v2)
		}
	}
	return sb.String()
}
