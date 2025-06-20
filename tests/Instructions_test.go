package tests_test

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestADC_ZX(t *testing.T) {
	th := ut.AssertHelperNew(t)
	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	tm := ut.DefaultTestMemory(t)

	data := []ut.InstructionTestData{
		ut.InstructionTestData{
			Name:                         "Normal Addition",
			AccumolatorSetup:             0x10,
			RegisterXSetup:               0,
			RegisterYSetup:               0,
			MemorySetup:                  []c.Word{c.Word(c.ADC_ZX), 0x00, 0x0F},
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x1F,
			ExpectedProcessorStatusValue: 0b00000000,
		},
		ut.InstructionTestData{
			Name:                         "Addition with Carry",
			AccumolatorSetup:             0x10,
			RegisterXSetup:               0,
			RegisterYSetup:               0,
			MemorySetup:                  []c.Word{c.Word(c.ADC_ZX), 0x00, 0xF1},
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x01,
			ExpectedProcessorStatusValue: 0b00000001,
		},
		ut.InstructionTestData{
			Name:                   "Addition with Zero result",
			AccumolatorSetup:       0x10,
			RegisterXSetup:         0,
			RegisterYSetup:         0,
			MemorySetup:            []c.Word{c.Word(c.ADC_ZX), 0x00, 0xF0},
			ExpectToAdvancedCycles: 4,
			ExpectAccumulatorValue: 0x00,
			// TODO check if carry is should be set?
			ExpectedProcessorStatusValue: 0b00000011,
		},
	}

	t.Logf("All tests for %s", c.ADC_ZX)
	for idx, testData := range data {
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name
		testData.Run(t, cpu, tm)
	}

	t.Run(" Test case 3: Overflow Case", func(t2 *testing.T) {
		t2.Skip("TODO: I don't understand Overflow mode yet. Its something with floating point arithmetics?")
		cpu.Accumulator = 0xFF
		cpu.RegisterX = 0x00
		cpu.Execute(1, tm, true)
		cpu.AssertCycle(12)
		th.AssertEqualsUint8(0x00, cpu.Accumulator, "Wrong Accumulator state")
		th.AssertEqualsUint8(0b10000000, cpu.Status, "Wrong Status register \n%v")
	})

	t.Run(" Test case 4: Negative Result", func(t2 *testing.T) {
		t2.Skip("TODO: I don't understand Overflow mode yet. Its something with floating point arithmetics?")
		cpu.Accumulator = c.Word(0x80)
		cpu.RegisterX = 0x00
		cpu.Execute(1, tm, true)
		th.AssertEqualsUint8(0x00, cpu.Accumulator, "Wrong Accumulator state")
		th.AssertEqualsUint8(0b10000010, cpu.Status, "Wrong Status register \nExpect: %#08b\n%v", 0b10000010)
		cpu.AssertCycle(20)
	})
}
