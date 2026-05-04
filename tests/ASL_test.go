package tests_test

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestASL(t *testing.T) {
	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	tm := ut.DefaultTestMemory(t)

	data := []ut.InstructionTestData{
		{
			Name:                         "ASL Accumulator",
			AccumolatorSetup:             0x05,
			RegisterXSetup:               0,
			RegisterYSetup:               0,
			MemorySetup:                  []any{c.ASL_A},
			ExpectToAdvancedCycles:       2,
			ExpectAccumulatorValue:       10,
			ExpectedProcessorStatusValue: 0b00000000,
		},
	}

	t.Logf("All tests for %s", c.ASL_A)
	for idx, testData := range data {
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name
		testData.Run(t, cpu, tm)
	}
}
