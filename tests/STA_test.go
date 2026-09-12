package tests

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestSTA(t *testing.T) {
	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	tm := ut.DefaultTestMemory(t)

	data := []ut.InstructionTestData{
		{
			Name:                         "STA Zero Page",
			AccumolatorSetup:             0x42,
			MemorySetup:                  []any{c.STA_Z, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, // Store at $001
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x4: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA Zero Page,X",
			AccumolatorSetup:             0x42,
			RegisterXSetup:               0x05,
			MemorySetup:                  []any{c.STA_ZX, 0x10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Store at $0010 + 5 = $0015
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x0015: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA Absolute",
			AccumolatorSetup:             0x42,
			MemorySetup:                  []any{c.STA_A, 0x34, 0x12}, // Store at $1234
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x1234: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA Absolute,X",
			AccumolatorSetup:             0x42,
			RegisterXSetup:               0x05,
			MemorySetup:                  []any{c.STA_AX, 0x30, 0x12}, // Store at $1230 + 5 = $1235
			ExpectToAdvancedCycles:       5,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x1235: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA Absolute,Y",
			AccumolatorSetup:             0x42,
			RegisterYSetup:               0x05,
			MemorySetup:                  []any{c.STA_AY, 0x30, 0x12}, // Store at $1230 + 5 = $1235
			ExpectToAdvancedCycles:       5,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x1235: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA (Indirect,X)",
			AccumolatorSetup:             0x42,
			RegisterXSetup:               0x05,
			MemorySetup:                  []any{c.STA_IX, 0x10, 0x20, 0x00}, // Pointer at $0010 points to $0020. Target: $0020 + 5 = $0025
			ExpectToAdvancedCycles:       6,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x0025: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "STA (Indirect),Y",
			AccumolatorSetup:             0x42,
			RegisterYSetup:               0x05,
			MemorySetup:                  []any{c.STA_IY, 0x10, 0x20, 0x00}, // Pointer at $0010 points to $0020. Target: $0020 + 5 = $0025
			ExpectToAdvancedCycles:       6,
			ExpectAccumulatorValue:       0x42,
			ExpectMemory:                 []any{0x0025: 0x42},
			ExpectedProcessorStatusValue: 0b00000000,
		},
	}

	t.Logf("All tests for STA")
	for idx, testData := range data {
		tm = ut.DefaultTestMemory(t)
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name
		testData.Run(t, cpu, tm)
	}
}
