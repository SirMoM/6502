package util_test

import (
	c "noah-ruben.com/6502/computer"
	"testing"
)

// TestHelper wraps *testing.T to provide assertion methods.
type AssertHelper struct {
	*testing.T
}

func AssertHelperNew(t *testing.T) AssertHelper { return AssertHelper{T: t} }

// AssertEquals checks if the expected value equals the actual value.
// It supports both uint8 and c.Word types for the actual value.
func (h *AssertHelper) AssertEqualsUint8(expected uint8, actual interface{}, message string, args ...any) {
	h.T.Helper()
	switch a := actual.(type) {
	case uint8:
		if expected != a {
			h.Fatalf(message, append(args, a)...)
		}
	case c.Word:
		if expected != uint8(a) {
			h.Fatalf(message, append(args, a)...)
		}
	case c.ProcessorStatus:
		if expected != uint8(a.Status) {
			h.Fatalf(message, append(args, a)...)
		}
	default:
		h.Fatalf("Unsupported type for actual value: %T", a)
	}
}

// CPU - Testing methods
type TestCpuLogger struct {
	t *testing.T
}

func NewTestCpuLogger(tee *testing.T) *TestCpuLogger {
	return &TestCpuLogger{tee}
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

// TestMemory - A READONLY continous memory without a fixed size.
type TestMemory struct {
	Data []c.Word
	t    *testing.T
	idx  uint
}

func DefaultTestMemory(tee *testing.T) *TestMemory {
	return NewTestMemory(tee, 100)
}

func NewTestMemory(tee *testing.T, size uint) *TestMemory {
	return &TestMemory{
		Data: make([]c.Word, 0, size),
		t:    tee,
		idx:  0,
	}

}

func (mem TestMemory) Init() error {
	mem.t.Helper()
	return nil
}
func (mem *TestMemory) WriteAddress(destination c.Address, address c.Address) {
	mem.t.Log("[WARN] DONT USE `WriteAddress(<args>)` it does nothing!")
}
func (mem *TestMemory) WriteWord(destination c.Address, value c.Word) {
	mem.t.Log("[WARN] DONT USE `WordWrite(<args>)` it does nothing!")
}
func (mem *TestMemory) ReadWord(source c.Address) c.Word {
	res := mem.Data[mem.idx]
	mem.idx++
	return res
}
func (mem *TestMemory) ReadAddress(source c.Address) c.Address {
	lsb := mem.ReadWord(source)
	msb := mem.ReadWord(source + 1)
	return c.Address(uint16(msb)<<8 | uint16(lsb))
}
func (mem *TestMemory) AppendInplace(data []c.Word) {
	for _, word := range data {
		mem.Data = append(mem.Data, word)
	}
}

// InstructionTestData can be used to create a Test from this "configuration"
type InstructionTestData struct {
	// Name - Name of  the Test
	Name string

	// AccumolatorSetup - Value that should be in the Accumolator before the computer executes the first Instruction
	AccumolatorSetup c.Word
	// RegisterXSetup - Value that should be in the X-Register before the computer executes the first Instruction
	RegisterXSetup c.Word
	// RegisterYSetup - Value that should be in the Y-Register before the computer executes the first Instruction
	RegisterYSetup c.Word
	// MemorySetup - Continous memory values. These will be appended to the TestMemory and will be "returned" FIFO as requests are made to the memory. This means memory jump operations are not exexuted. Just the next values are returned!
	MemorySetup []c.Word

	// ExpectToAdvancedCycles - The number of cycles that the compputer should have advanced
	ExpectToAdvancedCycles uint
	// ExpectAccumulatorValue - Values of the Accumulator after the computer execuest the Instruction
	ExpectAccumulatorValue uint8
	// ExpectedProcessorStatusValue - The Value of the [../../computer/status.go|ProcessorStatus] to have
	ExpectedProcessorStatusValue uint8
}

func (i InstructionTestData) Run(tee *testing.T, cpu *c.SixFiveOTwo, tm *TestMemory) {
	tee.Run(i.Name, func(t *testing.T) {
		assert := AssertHelperNew(t)
		currentCycles := cpu.Cycle
		// Setup
		cpu.Status.Reset()
		cpu.Accumulator = i.AccumolatorSetup
		cpu.RegisterX = i.RegisterXSetup
		cpu.RegisterY = i.RegisterYSetup

		tm.AppendInplace(i.MemorySetup)

		// Test
		cpu.Execute(1, tm, true)
		assert.AssertEqualsUint8(i.ExpectAccumulatorValue, cpu.Accumulator, "Wrong Accumulator state expected: %#02x but got: %v", i.ExpectAccumulatorValue)
		assert.AssertEqualsUint8(i.ExpectedProcessorStatusValue, cpu.Status, "Wrong Status register expected:\n        %#08b but got: \n%v", i.ExpectedProcessorStatusValue)
		cpu.AssertCycle(currentCycles + i.ExpectToAdvancedCycles)
	})

}
