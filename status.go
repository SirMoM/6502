package main

import (
	"fmt"
	"strings"
)

// Flag positions in the processor status register
const (
	CarryFlagPosition uint8 = iota
	ZeroFlagPosition
	InterruptDisableFlagPosition
	DecimalModeFlagPosition
	BreakCommandFlagPosition
	UNUSEDPosition
	OverflowFlagPosition
	NegativeFlagPosition
)

// ProcessorStatus represents the status register of the 6502 CPU
// As instructions are executed, a set of processor flags are set or clear to record the results of the operation.
// These flags and some additional control flags are held in a special status register. Each flag has a single bit within the register.
//
// Instructions exist to test the values of the various bits, to set or clear some of them and to push or pull the entire set to or from the stack.
type ProcessorStatus struct {
	// Carry Flag
	//
	// The carry flag is set if the last operation caused an overflow from bit 7 of the result or an underflow from bit 0.
	// This condition is set during arithmetic, comparison and during logical shifts.
	// It can be explicitly set using the 'Set Carry Flag' (SEC) instruction and cleared with 'Clear Carry Flag' (CLC).
	//
	// Zero Flag
	//
	// The zero flag is set if the result of the last operation as was zero.
	//
	// Interrupt Disable
	//
	// The interrupt disable flag is set if the program has executed a 'Set Interrupt Disable' (SEI) instruction.
	// While this flag is set the processor will not respond to interrupts from devices until it is cleared by a 'Clear Interrupt Disable' (CLI) instruction.
	//
	// Decimal Mode
	//
	// While the decimal mode flag is set the processor will obey the rules of Binary Coded Decimal (BCD) arithmetic during addition and subtraction.
	// The flag can be explicitly set using 'Set Decimal Flag' (SED) and cleared with 'Clear Decimal Flag' (CLD).
	//
	// Break Command
	//
	// The break command bit is set when a BRK instruction has been executed and an interrupt has been generated to process it.
	//
	// Overflow Flag
	//
	// The overflow flag is set during arithmetic operations if the result has yielded an invalid 2's complement result
	// (e.g. adding to positive numbers and ending up with a negative result: 64 + 64 => ─128).
	// It is determined by looking at the carry between bits 6 and 7 and between bit 7 and the carry flag.
	//
	// Negative Flag
	//
	// The negative flag is set if the result of the last operation had bit 7 set to a one.
	Status Word
}

// GetOverflowFlag returns the value of the overflow flag
func (ps *ProcessorStatus) GetOverflowFlag() Word {
	return ps.GetFlag(OverflowFlagPosition)
}

// SetOverflowFlag sets the overflow flag
func (ps *ProcessorStatus) SetOverflowFlag(b bool) {
	ps.SetFlag(bit6, b)
}

// GetInterruptDisableFlag returns the value of the interrupt disable flag
func (ps *ProcessorStatus) GetInterruptDisableFlag() Word {
	return ps.GetFlag(InterruptDisableFlagPosition)
}

// GetDecimalFlag returns the value of the decimal mode flag
func (ps *ProcessorStatus) GetDecimalFlag() Word {
	return ps.GetFlag(DecimalModeFlagPosition)
}

// SetDecimalFlag sets the decimal mode flag
func (ps *ProcessorStatus) SetDecimalFlag(b bool) {
	ps.SetFlag(bit3, b)
}

// SetInterruptDisableFlag sets the interrupt disable flag
func (ps *ProcessorStatus) SetInterruptDisableFlag(b bool) {
	ps.SetFlag(bit2, b)
}

// SetCarryFlag sets the carry flag
func (ps *ProcessorStatus) SetCarryFlag(b bool) {
	ps.SetFlag(bit0, b)
}

// GetCarryFlag returns the value of the carry flag
func (ps *ProcessorStatus) GetCarryFlag() Word {
	return ps.GetFlag(CarryFlagPosition)
}

// SetNegativeFlag sets the negative flag
func (ps *ProcessorStatus) SetNegativeFlag(b bool) {
	ps.SetFlag(bit7, b)
}

// GetNegativeFlag returns the value of the negative flag
func (ps *ProcessorStatus) GetNegativeFlag() Word {
	return ps.GetFlag(NegativeFlagPosition)
}

// SetZeroFlag sets the zero flag
func (ps *ProcessorStatus) SetZeroFlag(b bool) {
	ps.SetFlag(bit1, b)
}

// GetZeroFlag returns the value of the zero flag
func (ps *ProcessorStatus) GetZeroFlag() Word {
	return ps.GetFlag(ZeroFlagPosition)
}

// Reset initializes the status register to 0
func (ps *ProcessorStatus) Reset() {
	ps.Status = 0x00
}

// SetFlag sets or clears a flag in the status register
func (ps *ProcessorStatus) SetFlag(flag Word, set bool) {
	if set {
		ps.Status = ps.Status | flag
	} else {
		ps.Status = ps.Status &^ flag
	}
}

// GetFlag retrieves the value of the specified flag bit in the processor status register
func (ps *ProcessorStatus) GetFlag(flag uint8) Word {
	return ps.Status >> Word(flag) & 1
}

// String returns a string representation of the processor status
func (ps ProcessorStatus) String() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "PST  :  %#08b\n", ps.Status)
	fmt.Fprintf(&sb, "%9s┘│││││││\n", "Negative ─")
	fmt.Fprintf(&sb, "%11s┘││││││\n", "Overflow ─")
	fmt.Fprintf(&sb, "%12s┘│││││\n", " UNUSED ")
	fmt.Fprintf(&sb, "%13s│││││\n", "Break  ")
	fmt.Fprintf(&sb, "%13s┘││││\n", "Command ─")
	fmt.Fprintf(&sb, "%14s┘│││\n", "Decimal Mode ─")
	fmt.Fprintf(&sb, "%15s│││\n", "Interrupt  ")
	fmt.Fprintf(&sb, "%15s┘││\n", "Disable ─")
	fmt.Fprintf(&sb, "%16s┘│\n", "Zero ─")
	fmt.Fprintf(&sb, "%17s┘\n", "Carry ─")

	return sb.String()
}
