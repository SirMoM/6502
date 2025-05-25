package tests

import (
	c "noah-ruben.com/6502/computer"
	"testing"
)

type TestCpuLogger struct {
	t *testing.T
}

func (t TestCpuLogger) LogE(msg string, args ...any) {
	t.t.Logf(msg, args...)
}

func (t TestCpuLogger) LogS(msg string, args ...any) {
	t.t.Logf(msg, args...)
}

func (t TestCpuLogger) SetCycle(cycle uint) {
	t.t.Logf("%d:\n", cycle)
}

func (t TestCpuLogger) Close() error {
	return nil
}

func TestADC_ZX(t *testing.T) {
	cpu := c.NewSixFiveOTwo(TestCpuLogger{t})

	cpu.ProgramCounter = 0x0200

}
