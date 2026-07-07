package tests_test

import (
	"testing"

	c "noah-ruben.com/6502/computer"
	"noah-ruben.com/6502/programs"
)

func TestMult10Program(t *testing.T) {
	logger := c.SetupLogging()
	cpu := c.NewSixFiveOTwo(&logger)
	mem := c.Memory16K{}
	_ = mem.Init()

	cpu.Reset(&mem)
	cpu.Accumulator = 7

	_ = programs.Mult10Prog.CopyToMemory(cpu.ProgramCounter, &mem)

	cpu.Execute(12, &mem, true)

	if cpu.Accumulator != 70 {
		t.Fatalf("expected accumulator to be 70, got %d", cpu.Accumulator)
	}

	_ = logger.Close()
}
