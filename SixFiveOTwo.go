package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Word uint8
type Address uint16
type SixFiveOTwo struct {
	Cycle uint

	// ProgramCounter
	// The program counter is a 16 bit register which points to the next instruction to be executed. The value of program counter is modified automatically as instructions are executed.
	//
	// The value of the program counter can be modified by executing a jump, a relative branch or a subroutine call to another memory address or by returning from a subroutine or interrupt.
	ProgramCounter Address

	// StackPointer
	// The processor supports a 256 byte stack located between $0100 and $01FF. The stack pointer is an 8 bit register and holds the low 8 bits of the next free location on the stack. The location of the stack is fixed and cannot be moved.
	//
	// Pushing bytes to the stack causes the stack pointer to be decremented. Conversely pulling bytes causes it to be incremented.
	//
	// The CPU does not detect if the stack is overflowed by excessive pushing or pulling operations and will most likely result in the program crashing.
	StackPointer Word

	// Accumulator
	//
	// The 8 bit accumulator is used all arithmetic and logical operations (with the exception of increments and decrements). The contents of the accumulator can be stored and retrieved either from memory or the stack.
	//
	// Most complex operations will need to use the accumulator for arithmetic and efficient optimisation of its use is a key feature of time critical routines.
	Accumulator Word

	// Index Register X
	//
	// The 8 bit index register is most commonly used to hold counters or offsets for accessing memory. The value of the X register can be loaded and saved in memory, compared with values held in memory or incremented and decremented.
	//
	// The X register has one special function. It can be used to get a copy of the stack pointer or change its value.
	//
	// The Y register is similar to the X register in that it is available for holding counter or offsets memory access and supports the same set of memory load, save and compare operations as wells as increments and decrements.
	RegisterX, RegisterY Word

	Status ProcessorStatus
}

func (cpu SixFiveOTwo) String() string {
	sb := strings.Builder{}
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, " C   : %04d\n", cpu.Cycle)
	fmt.Fprintf(&sb, "PC   : %#04X\n", cpu.ProgramCounter)
	fmt.Fprintf(&sb, "SP   : %#02X\n", cpu.StackPointer)
	fmt.Fprintf(&sb, "ACC  : %#02X\n", cpu.Accumulator)
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, "REG X: %#02X\n", cpu.RegisterX)
	fmt.Fprintf(&sb, "REG Y: %#02X\n", cpu.RegisterY)
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, "%v\n", cpu.Status)
	return sb.String()
}
func (cpu *SixFiveOTwo) short() string {
	sb := strings.Builder{}
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, " C   : %04d\n", cpu.Cycle)
	fmt.Fprintf(&sb, "PC   : %#04X\n", cpu.ProgramCounter)
	fmt.Fprintf(&sb, "SP   : %#02X\n", cpu.StackPointer)
	fmt.Fprintf(&sb, "ACC  : %#02X\n", cpu.Accumulator)
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, "REG X: %#02X\n", cpu.RegisterX)
	fmt.Fprintf(&sb, "REG Y: %#02X\n", cpu.RegisterY)
	fmt.Fprintln(&sb, "────────────────")
	fmt.Fprintf(&sb, "Satus: %#02X\n", cpu.Status.Status)
	fmt.Fprintln(&sb, "────────────────")

	return sb.String()
}
func (cpu *SixFiveOTwo) Reset(mem *Memory) {
	mem.Init()

	cpu.ProgramCounter = 0xFFFC
	cpu.ProgramCounter = cpu.FetchAddress(mem)
	cpu.StackPointer = 0xFF
	cpu.Accumulator = 0
	cpu.RegisterX = 0
	cpu.RegisterY = 0
	cpu.Cycle = 0
	cpu.Status.Reset()

}

func (cpu *SixFiveOTwo) FetchInstruction(mem *Memory) Instruction {
	return Instruction(cpu.FetchWordFromProgrammCounter(mem))
}

// FetchAddress fetches a Address from Memmory and incremments the ProgrmmCounter
/*
	Memory Address | Contents       | Description
	---------------+----------------+------------------------------------
	$0200       |     $34        | LSB Least Significant Byte (Low Byte) of the address
	---------------+----------------+------------------------------------
	$0201       |     $12        | MSB Most Significant Byte (High Byte) of the address
	---------------+----------------+------------------------------------
*/
func (cpu *SixFiveOTwo) FetchAddress(mem *Memory) Address {
	lsb := cpu.FetchWordFromProgrammCounter(mem)
	msb := cpu.FetchWordFromProgrammCounter(mem)
	return Address(uint16(msb)<<8 | uint16(lsb))
}

func (cpu *SixFiveOTwo) FetchWord(mem *Memory, adress Address) Word {
	// fmt.Printf("Loading from: %#04X -> ", cpu.ProgramCounter)
	data := mem.Data[adress]
	cpu.Cycle++
	// fmt.Printf("%#02X\n", data)
	return data

}
func (cpu *SixFiveOTwo) FetchWordFromProgrammCounter(mem *Memory) Word {
	// fmt.Printf("Loading from: %#04X -> ", cpu.ProgramCounter)
	data := mem.Data[cpu.ProgramCounter]
	cpu.ProgramCounter++
	cpu.Cycle++
	// fmt.Printf("%#02X\n", data)
	return data

}

type Instruction uint8

const (
	// Load Accumulation
	LDA_I Instruction = 0xA9
	LDX_I Instruction = 0xA2

	// ADC - Add with Carry
	ADC_ZX Instruction = 0x75 // Zero Page,X

	// JMP - Jump
	JMP_ABS Instruction = 0x4C // Absolute
	JMP_IND Instruction = 0x6C // Indirect
)
const (
	bit7 Word = 0x80
	bit6 Word = 0x40
	bit5 Word = 0x20
	bit4 Word = 0x10
	bit3 Word = 0x8
	bit2 Word = 0x4
	bit1 Word = 0x2
	bit0 Word = 0x1
)
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

// evaluateAndSetStatusFlags updates the Zero (Z) and Negative (N) flags
// in the CPU's status register based on the provided data byte.
// The Zero flag is set if data is 0.
// The Negative flag is set if bit 7 (the most significant bit) of data is set.
// This is typically called after operations that affect these flags (e.g., loads,
// arithmetic, logic).
//
// Parameters:
//
//	data: The Word value to evaluate for setting the flags.
//
// Side Effects:
//
//	Modifies cpu.Status (specifically the Z and N flags).
func (cpu *SixFiveOTwo) evaluateAndSetStatusFlags(data Word) {
	zero := data == 0
	// Assumes NegativeFlagPosition is 7 for an 8-bit value.
	neg := (data >> NegativeFlagPosition) > 0
	cpu.Status.SetZeroFlag(zero)
	cpu.Status.SetNegativeFlag(neg)
}

// loadIntoRegister fetches a byte from memory using FetchWord (which advances the
// Program Counter) and stores it into the specified register.
// It then updates the Zero (Z) and Negative (N) status flags based on the value
// that was loaded, by calling evaluateAndSetStatusFlags.
// Parameters:
//
//	reg: A pointer to the target CPU register (e.g., &cpu.Accumulator, &cpu.RegisterX, &cpu.RegisterY).
//	mem: A pointer to the Memory interface/struct used for fetching the byte.
//
// Side Effects:
//   - Modifies the value of the register pointed to by `reg`.
//   - Increments the Program Counter (PC) via the call to FetchWord.
//   - Modifies cpu.Status (specifically the Z and N flags) via the call
//     to evaluateAndSetStatusFlags.
func (cpu *SixFiveOTwo) loadIntoRegister(reg *Word, mem *Memory) {
	data := cpu.FetchWordFromProgrammCounter(mem)
	*reg = data
	cpu.evaluateAndSetStatusFlags(data)
}

func (cpu *SixFiveOTwo) Execute(cyclesToRun uint, mem *Memory) {
	executionEnd := cpu.Cycle + cyclesToRun
	for cyclesToRun == 0 || cpu.Cycle <= executionEnd {
		fmt.Println(cpu.short())
		instruction := cpu.FetchInstruction(mem)
		switch instruction {
		case LDX_I:
			fmt.Println("LDX_I")
			cpu.loadIntoRegister(&cpu.RegisterX, mem)
		case LDA_I:
			fmt.Println("LDA_I")
			cpu.loadIntoRegister(&cpu.Accumulator, mem)
		case ADC_ZX:
			// 4 cycles
			fmt.Println("ADC_ZX")
			lhs := cpu.FetchWord(mem, Address(cpu.RegisterX))
			res := lhs + cpu.Accumulator
			println(res)
		case JMP_ABS:
			fmt.Printf("cc: %d", cpu.Cycle)
			cpu.ProgramCounter = cpu.FetchAddress(mem)
			cpu.Cycle++

		default:
			fmt.Printf("INSTRUCTION: %#04X\n", instruction)
			fmt.Fprintln(os.Stderr, cpu)
			os.Exit(1)
		}

	}

}

// Processor Status
// As instructions are executed a set of processor flags are set or clear to record the results of the operation. This flags and some additional control flags are held in a special status register. Each flag has a single bit within the register.
//
// Instructions exist to test the values of the various bits, to set or clear some of them and to push or pull the entire set to or from the stack.
type ProcessorStatus struct {

	// Carry Flag
	//
	// The carry flag is set if the last operation caused an overflow from bit 7 of the result or an underflow from bit 0. This condition is set during arithmetic, comparison and during logical shifts. It can be explicitly set using the 'Set Carry Flag' (SEC) instruction and cleared with 'Clear Carry Flag' (CLC).
	//
	// Zero Flag
	//
	// The zero flag is set if the result of the last operation as was zero.
	//
	// Interrupt Disable
	//
	// The interrupt disable flag is set if the program has executed a 'Set Interrupt Disable' (SEI) instruction. While this flag is set the processor will not respond to interrupts from devices until it is cleared by a 'Clear Interrupt Disable' (CLI) instruction.
	//
	// Decimal Mode
	//
	// While the decimal mode flag is set the processor will obey the rules of Binary Coded Decimal (BCD) arithmetic during addition and subtraction. The flag can be explicitly set using 'Set Decimal Flag' (SED) and cleared with 'Clear Decimal Flag' (CLD).
	//
	// Break Command
	//
	// The break command bit is set when a BRK instruction has been executed and an interrupt has been generated to process it.
	//
	// Overflow Flag
	//
	// The overflow flag is set during arithmetic operations if the result has yielded an invalid 2's complement result (e.g. adding to positive numbers and ending up with a negative result: 64 + 64 => ─128). It is determined by looking at the carry between bits 6 and 7 and between bit 7 and the carry flag.
	//
	// Negative Flag
	//
	// The negative flag is set if the result of the last operation had bit 7 set to a one.
	Status Word
}

func (ps *ProcessorStatus) GetOverflowFlag() Word {
	return ps.GetFlag(OverflowFlagPosition)
}
func (ps *ProcessorStatus) SetOverflowFlag(b bool) {
	ps.SetFlag(bit6, b)
}

func (ps *ProcessorStatus) GetInterruptDisableFlag() Word {
	return ps.GetFlag(InterruptDisableFlagPosition)
}

func (ps *ProcessorStatus) GetDecimalFlag() Word {
	return ps.GetFlag(DecimalModeFlagPosition)
}
func (ps *ProcessorStatus) SetDecimalFlag(b bool) {
	ps.SetFlag(bit3, b)
}

func (ps *ProcessorStatus) SetInterruptDisableFlag(b bool) {
	ps.SetFlag(bit2, b)
}

func (ps *ProcessorStatus) SetCarryFlag(b bool) {
	ps.SetFlag(bit0, b)
}
func (ps *ProcessorStatus) GetCarryFlag() Word {
	return ps.GetFlag(CarryFlagPosition)
}

func (ps *ProcessorStatus) SetNegativeFlag(b bool) {
	ps.SetFlag(bit7, b)
}
func (ps *ProcessorStatus) GetNegativeFlag() Word {
	return ps.GetFlag(NegativeFlagPosition)
}

func (ps *ProcessorStatus) SetZeroFlag(b bool) {
	ps.SetFlag(bit1, b)
}
func (ps *ProcessorStatus) GetZeroFlag() Word {
	return ps.GetFlag(ZeroFlagPosition)
}

func (ps *ProcessorStatus) Reset() {
	ps.Status = 0x00

}

func (ps *ProcessorStatus) SetFlag(flag Word, set bool) {
	if set {
		ps.Status = ps.Status | flag
	} else {
		ps.Status = ps.Status &^ flag
	}
}

// GetFlag retrieves the value of the specified flag bit in the processor status register.
func (ps *ProcessorStatus) GetFlag(flag uint8) Word {
	return ps.Status >> Word(flag) & 1
}

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

type Memory struct {
	Data []Word
}

func (mem *Memory) Init() {
	mem.Data = make([]Word, math.MaxUint16+1)
	mem.Data[0xFFFC] = 0x00
	mem.Data[0xFFFD] = 0x02

	mem.Data[0x0200] = 0xF9
	mem.Data[0x0201] = Word(LDX_I)
	mem.Data[0x0202] = 0x0F
	mem.Data[0x0203] = Word(ADC_ZX)

}

func (mem Memory) String() string {
	sb := strings.Builder{}
	blockSize := 32
	collums := 2
	for i := 0; i < math.MaxUint16; i += blockSize {
		fmt.Fprintf(&sb, "── Blocks from %#04X to %#04X ──\n", i, i+blockSize-1)
		for j := 0; j < blockSize/collums; j++ {
			i1 := i + j
			v1 := mem.Data[i1]
			fmt.Fprintf(&sb, "[%#04X] %#02X ─ ", i1, v1)
			i2 := j + i + blockSize/collums
			v2 := mem.Data[i2]
			fmt.Fprintf(&sb, "[%#04X] %#02X\n", i2, v2)
		}
	}
	return sb.String()
}
