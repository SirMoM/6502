package tests_test

import (
	c "noah-ruben.com/6502/computer"
	"noah-ruben.com/6502/programs"
	"testing"
)

var gt *testing.T
var ps *c.ProcessorStatus
var cleanupLoggerFiles func()

func assertEquals(expected uint8, actual c.Word, message string, args ...any) {
	if expected != uint8(actual) {
		gt.Fatalf(message, append(args, actual)...)
	}
}

func ZeroFlag(t *testing.T) {
	ps.SetZeroFlag(true)
	assertEquals(0b00000010, ps.Status, "Status should be 0b00000010 but was %d")

	ps.SetZeroFlag(true)
	assertEquals(0b00000010, ps.Status, "Status should be 0b00000010 but was %d")
	res := ps.GetZeroFlag()
	assertEquals(1, res, "Zero flag should be 1 but was %d")

	ps.SetZeroFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetZeroFlag()
	assertEquals(0, res, "Zero flag should be 0 but was %d")
	ps.SetZeroFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetZeroFlag()
	assertEquals(0, res, "Zero flag should be 0 but was %d")
}

func CarryFlag(t *testing.T) {
	ps.SetCarryFlag(true)
	assertEquals(0b00000001, ps.Status, "Status should be 0b00000001 but was %d")

	ps.SetCarryFlag(true)
	assertEquals(0b00000001, ps.Status, "Status should be 0b00000001 but was %d")
	res := ps.GetCarryFlag()
	assertEquals(1, res, "Carry flag should be 1 but was %d")

	ps.SetCarryFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetCarryFlag()
	assertEquals(0, res, "Carry flag should be 0 but was %d")
}

func NegativeFlag(t *testing.T) {
	ps.SetNegativeFlag(true)
	assertEquals(0b10000000, ps.Status, "Status should be 0b10000000 but was %d")

	ps.SetNegativeFlag(true)
	assertEquals(0b10000000, ps.Status, "Status should be 0b10000000 but was %d")
	res := ps.GetNegativeFlag()
	assertEquals(1, res, "Negative flag should be 1 but was %d")

	ps.SetNegativeFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetNegativeFlag()
	assertEquals(0, res, "Negative flag should be 0 but was %d")
}

func InterruptDisableFlag(t *testing.T) {
	ps.SetInterruptDisableFlag(true)
	assertEquals(0b00000100, ps.Status, "Status should be 0b00000100 but was %d")

	ps.SetInterruptDisableFlag(true)
	assertEquals(0b00000100, ps.Status, "Status should be 0b00000100 but was %d")
	res := ps.GetInterruptDisableFlag()
	assertEquals(1, res, "Interrupt Disable flag should be 1 but was %d")

	ps.SetInterruptDisableFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetInterruptDisableFlag()
	assertEquals(0, res, "Interrupt Disable flag should be 0 but was %d")
}

func DecimalFlag(t *testing.T) {
	ps.SetDecimalFlag(true)
	assertEquals(0b00001000, ps.Status, "Status should be 0b00001000 but was %d")

	ps.SetDecimalFlag(true)
	assertEquals(0b00001000, ps.Status, "Status should be 0b00001000 but was %d")
	res := ps.GetDecimalFlag()
	assertEquals(1, res, "Decimal flag should be 1 but was %d")

	ps.SetDecimalFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetDecimalFlag()
	assertEquals(0, res, "Decimal flag should be 0 but was %d")
}

func OverflowFlag(t *testing.T) {
	ps.SetOverflowFlag(true)
	assertEquals(0b01000000, ps.Status, "Status should be 0b01000000 but was %d")

	ps.SetOverflowFlag(true)
	assertEquals(0b01000000, ps.Status, "Status should be 0b01000000 but was %d")
	res := ps.GetOverflowFlag()
	assertEquals(1, res, "Overflow flag should be 1 but was %d")

	ps.SetOverflowFlag(false)
	assertEquals(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetOverflowFlag()
	assertEquals(0, res, "Overflow flag should be 0 but was %d")
}

func TestStatusRegister(t *testing.T) {
	gt = t
	tmp := c.ProcessorStatus{}
	ps = &tmp
	ps.Reset()
	gt.Log(ps)
	cleanupLoggerFiles = c.SetupLogging()
	defer cleanupLoggerFiles()

	t.Run("Test Zero Flag behavior", ZeroFlag)
	t.Run("Test Carry Flag behavior", CarryFlag)
	t.Run("Test Interrupt Disable Flag behavior", InterruptDisableFlag)
	t.Run("Test Decimal Flag behavior", DecimalFlag)
	t.Run("Test Overflow Flag behavior", OverflowFlag)
	t.Run("Test Negative Flag behavior", NegativeFlag)

	gt.Log(ps)

}

func TestMiniProgramm(t *testing.T) {
	cleanupLoggerFiles = c.SetupLogging()
	defer cleanupLoggerFiles()
	gt = t
	cpu := c.SixFiveOTwo{}
	mem := c.Memory{}
	mem.Init()
	cpu.Reset(&mem)

	_ = programs.MiniProg.CopyToMemory(cpu.ProgramCounter, &mem)

	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(2)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(4)
	//cpu.Execute(1, &mem, true)
	//cpu.AssertCycle(8)
	t.Log(cpu)
}
