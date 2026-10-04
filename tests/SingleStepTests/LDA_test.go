package tests

import (
	"encoding/json"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestExample(t *testing.T) {
	tests := *LoadFile()
	test := tests[0]
	b, err := json.MarshalIndent(test, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	cpu.ProgramCounter = test.Initial.PC
	cpu.Accumulator = test.Initial.A
	cpu.RegisterX = test.Initial.X
	cpu.RegisterY = test.Initial.Y
	cpu.StackPointer = test.Initial.S
	cpu.Status = test.Initial.P
}
