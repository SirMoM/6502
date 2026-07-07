package tests

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestCLC(t *testing.T) {
	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	tm := ut.DefaultTestMemory(t)

	data := []ut.InstructionTestData{
		{
			Name:                         "CLC Implied",
			MemorySetup:                  []any{c.CLC},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b00000001},
			ExpectToAdvancedCycles:       2,
			ExpectAccumulatorValue:       0x0,        // Not affected
			RegisterXSetup:               0x0,        // Not affected
			RegisterYSetup:               0x0,        // Not affected
			ExpectedProcessorStatusValue: 0b00000000, // Carry flag (C) is set to 0, other flags remain unchanged
		},
		{
			Name:                         "CLC Implied",
			MemorySetup:                  []any{c.CLC},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b11111111},
			ExpectToAdvancedCycles:       2,
			ExpectAccumulatorValue:       0x0,        // Not affected
			RegisterXSetup:               0x0,        // Not affected
			RegisterYSetup:               0x0,        // Not affected
			ExpectedProcessorStatusValue: 0b11111110, // Carry flag (C) is set to 0, other flags remain unchanged
		},
	}

	t.Logf("All tests for CLC")
	for idx, testData := range data {
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name
		testData.Run(t, cpu, tm)
	}
}
