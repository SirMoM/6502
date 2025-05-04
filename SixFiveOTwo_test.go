package main_test

import (
	"fmt"
	"testing"

	m "noah-ruben.com/6502"
)

var gt *testing.T

func assertEquals(expected uint8, actual uint8, message string, args ...any) {
	if expected != actual {
		gt.Fatalf(message, args)
	}

}

func TestStatusRegister(t *testing.T) {
	gt = t
	ps := m.ProcessorStatus{}
	ps.Reset()

	ps.SetZeroFlag(true)

	assertEquals(2, ps.Status, "Status should be 0b00000010 but was %d", ps.Status)

	fmt.Println(ps)
}
