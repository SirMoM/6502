package programs

import (
	c "noah-ruben.com/6502/computer"
)

type Program interface {
	CopyToMemory(mem *c.Memory) error
}
