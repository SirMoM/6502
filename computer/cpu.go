package computer

import (
	"fmt"
	"os"
	"strings"
)

type Word uint8

func (w Word) String() string {
	return fmt.Sprintf("%#02x", uint8(w))
}

type Address uint16

func (a Address) String() string {
	return fmt.Sprintf("%#04x", uint16(a))
}

// SixFiveOTwo represents the 6502 CPU
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

	logger CpuLogger
}

func NewSixFiveOTwo(logger CpuLogger) *SixFiveOTwo {
	return &SixFiveOTwo{
		logger: logger,
	}
}

// String returns a string representation of the CPU state
func (cpu SixFiveOTwo) String() string {
	sb := strings.Builder{}
	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, " C   : %04d\n", cpu.Cycle)
	_, _ = fmt.Fprintf(&sb, "PC   : %s\n", cpu.ProgramCounter)
	_, _ = fmt.Fprintf(&sb, "SP   : %s\n", cpu.StackPointer)
	_, _ = fmt.Fprintf(&sb, "ACC  : %s\n", cpu.Accumulator)
	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, "REG X: %s\n", cpu.RegisterX)
	_, _ = fmt.Fprintf(&sb, "REG Y: %s\n", cpu.RegisterY)
	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, "%v\n", cpu.Status)
	return sb.String()
}

// Short returns a compact string representation of the CPU state
func (cpu *SixFiveOTwo) Short() string {
	sb := strings.Builder{}

	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, " C   : %04d\n", cpu.Cycle)
	_, _ = fmt.Fprintf(&sb, "PC   : %s\n", cpu.ProgramCounter)
	_, _ = fmt.Fprintf(&sb, "SP   : %s\n", cpu.StackPointer)
	_, _ = fmt.Fprintf(&sb, "ACC  : %s\n", cpu.Accumulator)
	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, "REG X: %s\n", cpu.RegisterX)
	_, _ = fmt.Fprintf(&sb, "REG Y: %s\n", cpu.RegisterY)
	_, _ = fmt.Fprintln(&sb, "────────────────")
	_, _ = fmt.Fprintf(&sb, "Satus: %s\n", cpu.Status.Status)
	_, _ = fmt.Fprintln(&sb, "────────────────")

	return sb.String()
}

// Reset initializes the CPU to its initial state
func (cpu *SixFiveOTwo) Reset(mem Memory) {
	_ = mem.Init()
	cpu.ProgramCounter = 0xFFFC
	cpu.StackPointer = 0xFF
	cpu.Accumulator = 0
	cpu.RegisterX = 0
	cpu.RegisterY = 0
	cpu.Status.Reset()
	cpu.Cycle = 0

}

func (cpu *SixFiveOTwo) addCycle() {
	cpu.Cycle = cpu.Cycle + 1
	cpu.logger.SetCycle(cpu.Cycle)
}

// FetchInstruction fetches the next instruction from memory
func (cpu *SixFiveOTwo) FetchInstruction(mem Memory) Instruction {
	return Instruction(cpu.FetchWordFromProgramCounter(mem))
}

// FetchAddress fetches a Address from Memory and increments the ProgramCounter
/*
	Memory Address | Contents       | Description
	---------------+----------------+------------------------------------
	$0200       |     $34        | LSB Least Significant Byte (Low Byte) of the address
	---------------+----------------+------------------------------------
	$0201       |     $12        | MSB Most Significant Byte (High Byte) of the address
	---------------+----------------+------------------------------------
*/
func (cpu *SixFiveOTwo) FetchAddress(mem Memory) Address {
	addr := mem.ReadAddress(cpu.ProgramCounter)
	cpu.addCycle()
	cpu.ProgramCounter++
	cpu.addCycle()
	cpu.ProgramCounter++
	return addr
}

// FetchWord fetches a Word from Memory at the specified address
func (cpu *SixFiveOTwo) FetchWord(mem Memory, address Address) Word {
	data := mem.ReadWord(address)
	cpu.addCycle()
	return data
}

// FetchWordFromProgramCounter fetches a Word from Memory at the ProgramCounter and increments it
func (cpu *SixFiveOTwo) FetchWordFromProgramCounter(mem Memory) Word {
	data := mem.ReadWord(cpu.ProgramCounter)
	cpu.ProgramCounter++
	cpu.addCycle()
	return data
}

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
func (cpu *SixFiveOTwo) loadIntoRegister(reg *Word, mem Memory) {
	data := cpu.FetchWordFromProgramCounter(mem)
	*reg = data
	cpu.evaluateAndSetStatusFlags(data)
}

// Execute runs the CPU for the specified number of cycles
func (cpu *SixFiveOTwo) Execute(cyclesToRun uint, mem Memory, verbose bool) {
	executionEnd := cpu.Cycle + cyclesToRun
	for cyclesToRun == 0 || cpu.Cycle <= executionEnd {
		instruction := cpu.FetchInstruction(mem)
		cpu.logger.LogE("%s\n", instruction)
		switch instruction {
		case LDX_I:
			cpu.loadIntoRegister(&cpu.RegisterX, mem)
			cpu.logger.LogE("%s\n", cpu.RegisterX)
		case LDA_I:
			cpu.loadIntoRegister(&cpu.Accumulator, mem)
			cpu.logger.LogE("%s\n", cpu.Accumulator)
		case ADC_ZX:
			// 4 cycles
			// todo: finish
			lhs := cpu.FetchWord(mem, Address(cpu.RegisterX))
			res := lhs + cpu.Accumulator
			println(res)
			fmt.Printf(" %s + %s = %s", cpu.Accumulator, cpu.RegisterX, res)
			cpu.logger.LogE("%s", cpu.Accumulator)
		case JMP_ABS:
			fmt.Printf("cc: %d", cpu.Cycle)
			cpu.ProgramCounter = cpu.FetchAddress(mem)
			cpu.addCycle()
			cpu.logger.LogE("%s", cpu.ProgramCounter)

		case JMP_IND:
			fmt.Printf("cc: %d", cpu.Cycle)
			cpu.ProgramCounter = cpu.FetchAddress(mem)
			cpu.ProgramCounter = cpu.FetchAddress(mem)
			cpu.addCycle()
			cpu.logger.LogE("%s", cpu.ProgramCounter.String())

		default:
			cpu.logger.LogE("\n===============\n")
			cpu.logger.LogE("CPU CRASHED\n")
			cpu.logger.LogE("%s\n", instruction)
			os.Exit(1)
		}
	}
}

func (cpu SixFiveOTwo) AssertCycle(cycle uint) {
	if cpu.Cycle != cycle {
		_, _ = fmt.Fprintf(os.Stderr, "CPU is in the wrong cycle %d expected %d", cpu.Cycle, cycle)
		os.Exit(-1)
	}
}
