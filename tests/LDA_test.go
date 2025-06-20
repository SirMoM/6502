package tests_test

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestLDA(t *testing.T) {
	cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
	tm := ut.DefaultTestMemory(t)

	data := []ut.InstructionTestData{
		ut.InstructionTestData{
			Name:                         "LDA Load Immediate",
			AccumolatorSetup:             0x00,
			RegisterXSetup:               0xFF,
			RegisterYSetup:               0,
			MemorySetup:                  []c.Word{c.Word(c.LDA_I), 0xFF},
			ExpectToAdvancedCycles:       2,
			ExpectAccumulatorValue:       0xFF,
			ExpectedProcessorStatusValue: 0b10000000,
		},
		ut.InstructionTestData{
			Name:                         "LDA Load From Zero Page",
			AccumolatorSetup:             0x00,
			RegisterXSetup:               0xFF,
			RegisterYSetup:               0,
			MemorySetup:                  []c.Word{c.Word(c.LDA_I), 0xFF},
			ExpectToAdvancedCycles:       2,
			ExpectAccumulatorValue:       0xFF,
			ExpectedProcessorStatusValue: 0b10000000,
		},
	}

	t.Logf("All tests for %s", c.LDA_I)
	for idx, testData := range data {
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name
		testData.Run(t, cpu, tm)
	}
}
