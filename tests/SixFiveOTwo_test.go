package tests_test

import (
	c "noah-ruben.com/6502/computer"
	"noah-ruben.com/6502/programs"
	tu "noah-ruben.com/6502/tests/util"
	"testing"
)

var ah tu.AssertHelper
var ps *c.ProcessorStatus

func ZeroFlag(t *testing.T) {
	ps.SetZeroFlag(true)
	ah.AssertEqualsUint8(0b00000010, ps.Status, "Status should be 0b00000010 but was %d")

	ps.SetZeroFlag(true)
	ah.AssertEqualsUint8(0b00000010, ps.Status, "Status should be 0b00000010 but was %d")
	res := ps.GetZeroFlag()
	ah.AssertEqualsUint8(1, res, "Zero flag should be 1 but was %d")

	ps.SetZeroFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetZeroFlag()
	ah.AssertEqualsUint8(0, res, "Zero flag should be 0 but was %d")
	ps.SetZeroFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetZeroFlag()
	ah.AssertEqualsUint8(0, res, "Zero flag should be 0 but was %d")
}

func CarryFlag(t *testing.T) {
	ps.SetCarryFlag(true)
	ah.AssertEqualsUint8(0b00000001, ps.Status, "Status should be 0b00000001 but was %d")

	ps.SetCarryFlag(true)
	ah.AssertEqualsUint8(0b00000001, ps.Status, "Status should be 0b00000001 but was %d")
	res := ps.GetCarryFlag()
	ah.AssertEqualsUint8(1, res, "Carry flag should be 1 but was %d")

	ps.SetCarryFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetCarryFlag()
	ah.AssertEqualsUint8(0, res, "Carry flag should be 0 but was %d")
}

func NegativeFlag(t *testing.T) {
	ps.SetNegativeFlag(true)
	ah.AssertEqualsUint8(0b10000000, ps.Status, "Status should be 0b10000000 but was %d")

	ps.SetNegativeFlag(true)
	ah.AssertEqualsUint8(0b10000000, ps.Status, "Status should be 0b10000000 but was %d")
	res := ps.GetNegativeFlag()
	ah.AssertEqualsUint8(1, res, "Negative flag should be 1 but was %d")

	ps.SetNegativeFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetNegativeFlag()
	ah.AssertEqualsUint8(0, res, "Negative flag should be 0 but was %d")
}

func InterruptDisableFlag(t *testing.T) {
	ps.SetInterruptDisableFlag(true)
	ah.AssertEqualsUint8(0b00000100, ps.Status, "Status should be 0b00000100 but was %d")

	ps.SetInterruptDisableFlag(true)
	ah.AssertEqualsUint8(0b00000100, ps.Status, "Status should be 0b00000100 but was %d")
	res := ps.GetInterruptDisableFlag()
	ah.AssertEqualsUint8(1, res, "Interrupt Disable flag should be 1 but was %d")

	ps.SetInterruptDisableFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetInterruptDisableFlag()
	ah.AssertEqualsUint8(0, res, "Interrupt Disable flag should be 0 but was %d")
}

func DecimalFlag(t *testing.T) {
	ps.SetDecimalFlag(true)
	ah.AssertEqualsUint8(0b00001000, ps.Status, "Status should be 0b00001000 but was %d")

	ps.SetDecimalFlag(true)
	ah.AssertEqualsUint8(0b00001000, ps.Status, "Status should be 0b00001000 but was %d")
	res := ps.GetDecimalFlag()
	ah.AssertEqualsUint8(1, res, "Decimal flag should be 1 but was %d")

	ps.SetDecimalFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetDecimalFlag()
	ah.AssertEqualsUint8(0, res, "Decimal flag should be 0 but was %d")
}

func OverflowFlag(t *testing.T) {
	ps.SetOverflowFlag(true)
	ah.AssertEqualsUint8(0b01000000, ps.Status, "Status should be 0b01000000 but was %d")

	ps.SetOverflowFlag(true)
	ah.AssertEqualsUint8(0b01000000, ps.Status, "Status should be 0b01000000 but was %d")
	res := ps.GetOverflowFlag()
	ah.AssertEqualsUint8(1, res, "Overflow flag should be 1 but was %d")

	ps.SetOverflowFlag(false)
	ah.AssertEqualsUint8(0b00000000, ps.Status, "Status should be 0b00000000 but was %d")
	res = ps.GetOverflowFlag()
	ah.AssertEqualsUint8(0, res, "Overflow flag should be 0 but was %d")
}

func TestStatusRegister(t *testing.T) {
	ah = tu.AssertHelperNew(t)
	tmp := c.ProcessorStatus{}
	ps = &tmp
	ps.Reset()
	t.Log(ps)

	t.Run("Test Zero Flag behavior", ZeroFlag)
	t.Run("Test Carry Flag behavior", CarryFlag)
	t.Run("Test Interrupt Disable Flag behavior", InterruptDisableFlag)
	t.Run("Test Decimal Flag behavior", DecimalFlag)
	t.Run("Test Overflow Flag behavior", OverflowFlag)
	t.Run("Test Negative Flag behavior", NegativeFlag)

	t.Log(ps)

}

func TestMiniProgramm(t *testing.T) {
	gt = t
	logger := c.SetupLogging()

	cpu := c.NewSixFiveOTwo(&logger)

	mem := c.Memory16K{}
	_ = mem.Init()

	cpu.Reset(&mem)

	_ = programs.MiniProg.CopyToMemory(cpu.ProgramCounter, &mem)

	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(2)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(4)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(8)
	t.Log(cpu)

	_ = logger.Close()
}
