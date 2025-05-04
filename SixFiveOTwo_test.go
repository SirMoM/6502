package main_test

import (
	"fmt"
	"testing"

	m "noah-ruben.com/6502"
)

var gt *testing.T
var ps *m.ProcessorStatus

func assertEquals(expected uint8, actual uint8, message string, args ...any) {
	if expected != actual {
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
	tmp := m.ProcessorStatus{}
	ps = &tmp
	ps.Reset()
	gt.Log(ps)

	t.Run("Test Zero Flag behavior", ZeroFlag)
	t.Run("Test Carry Flag behavior", CarryFlag)
	t.Run("Test Interrupt Disable Flag behavior", InterruptDisableFlag)
	t.Run("Test Decimal Flag behavior", DecimalFlag)
	t.Run("Test Overflow Flag behavior", OverflowFlag)
	t.Run("Test Negative Flag behavior", NegativeFlag)

	gt.Log(ps)

}

func TestAsd(t *testing.T) {
	gt = t
	tmp := m.ProcessorStatus{}
	ps = &tmp
	ps.Reset()
	data := 0b11111001
	zero := data == 0
	fmt.Printf("z %t\n", zero)
	ps.SetZeroFlag(zero)
	fmt.Println(ps)

	// t.FailNow()
	// neg := (data >> 7) > 0
	// fmt.Printf("n %t\n", neg)
	// ps.SetNegativeFlag(neg)
	// fmt.Printf("%#08b\n", data)
	// fmt.Println(ps)
}
