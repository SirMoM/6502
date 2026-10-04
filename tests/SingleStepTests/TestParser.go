package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	c "noah-ruben.com/6502/computer"
)

func LoadFile() *[]SingleStepTest {
	data, err := os.ReadFile("bd.json")
	if err != nil {
		panic(err)
	}

	var tests []SingleStepTest
	if err := json.Unmarshal(data, &tests); err != nil {
		log.Fatal(err)
	}

	return &tests
}

// SingleStepTest is one test case from SingleStepTests/65x02 (e.g. 6502/v1/bd.json).
// Each file is a JSON array of these, so it decodes into a []SingleStepTest.
type SingleStepTest struct {
	Name    string     `json:"name"`
	Initial CPUState   `json:"initial"`
	Final   CPUState   `json:"final"`
	Cycles  []BusCycle `json:"cycles"`
}

// CPUState is the register file plus sparse RAM before or after the instruction.
type CPUState struct {
	PC  c.Address  `json:"pc"`
	S   c.Word     `json:"s"`
	A   c.Word     `json:"a"`
	X   c.Word     `json:"x"`
	Y   c.Word     `json:"y"`
	P   c.Word     `json:"p"`
	RAM []RAMEntry `json:"ram"`
}

// RAMEntry is one memory cell. JSON shape: [address, value]
type RAMEntry struct {
	Address c.Address
	Value   c.Word
}

// BusOp is the direction of one bus access.
type BusOp string

const (
	BusRead  BusOp = "read"
	BusWrite BusOp = "write"
)

// BusCycle is the single bus access made during one clock cycle.
// JSON shape: [address, value, "read" | "write"]
type BusCycle struct {
	Address c.Address
	Value   c.Word
	Op      BusOp
}

func (e *RAMEntry) UnmarshalJSON(data []byte) error {
	return unmarshalTuple(data, &e.Address, &e.Value)
}

func (b *BusCycle) UnmarshalJSON(data []byte) error {
	if err := unmarshalTuple(data, &b.Address, &b.Value, &b.Op); err != nil {
		return err
	}
	if b.Op != BusRead && b.Op != BusWrite {
		return fmt.Errorf("unknown bus op %q in %s", b.Op, data)
	}
	return nil
}

// unmarshalTuple decodes a JSON array such as [15099, 189, "read"]
// element by element into the given pointers, in order.
func unmarshalTuple(data []byte, fields ...any) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if len(raw) != len(fields) {
		return fmt.Errorf("expected %d elements, got %d: %s", len(fields), len(raw), data)
	}
	for i, field := range fields {
		if err := json.Unmarshal(raw[i], field); err != nil {
			return fmt.Errorf("element %d of %s: %w", i, data, err)
		}
	}
	return nil
}
