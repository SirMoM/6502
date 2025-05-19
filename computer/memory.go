package computer

import (
	"fmt"
	"math"
	"strings"
)

type Memory interface {
	Init() error
	WriteAddress(destination Address, address Address)
	WriteWord(destination Address, value Word)
	ReadWord(source Address) Word

	// ReadAddress fetches a Address from Memory and increments the ProgramCounter
	/*
		Memory Address | Contents       | Description
		---------------+----------------+------------------------------------
		$0200       |     $34        | LSB Least Significant Byte (Low Byte) of the address
		---------------+----------------+------------------------------------
		$0201       |     $12        | MSB Most Significant Byte (High Byte) of the address
		---------------+----------------+------------------------------------
	*/
	ReadAddress(source Address) Address
}

type Memory16K struct {
	Data []Word
}

func (mem *Memory16K) WriteWord(destination Address, value Word) {
	mem.Data[destination] = value
}

func (mem *Memory16K) ReadWord(source Address) Word {
	return mem.Data[source]
}

func (mem *Memory16K) ReadAddress(source Address) Address {
	lsb := mem.ReadWord(source)
	msb := mem.ReadWord(source + 1)
	return Address(uint16(msb)<<8 | uint16(lsb))
}

// Init initializes the memory with default values
func (mem *Memory16K) Init() error {
	mem.Data = make([]Word, math.MaxUint16+1)
	return nil
}

func (mem *Memory16K) WriteAddress(destination Address, address Address) {
	lsb := Word(address)
	msb := Word(address >> 8)
	mem.Data[destination] = lsb
	mem.Data[destination+1] = msb
}

// String returns a string representation of the memory
func (mem Memory16K) String() string {
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
