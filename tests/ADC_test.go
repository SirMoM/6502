package tests

import (
	"strconv"
	"testing"

	c "noah-ruben.com/6502/computer"
	ut "noah-ruben.com/6502/tests/util"
)

func TestADC(t *testing.T) {
	data := []ut.InstructionTestData{
		//{
		//	Name: "ADC Immediate",
		//	AccumolatorSetup: 0x42,
		//	MemorySetup: []any{c.ADC_I, 0x05},
		//	ExpectToAdvancedCycles: 2,
		//	ExpectAccumulatorValue: 0x47, // 0x42 + 0x05 = 0x47, no carry
		//	ExpectedProcessorStatusValue: 0b00000000, // No overflow, no zero
		//},
		//{
		//	Name: "ADC Immediate with Carry",
		//	AccumolatorSetup: 0x42,
		//	MemorySetup: []any{c.ADC_I 0x05},
		//	ProcessorStatusValueSetup: 0b10000000, // Carry flag set
		//	ExpectToAdvancedCycles: 2,
		//	ExpectAccumulatorValue: 0x47, // 0x42 + 0x05 + 0x01 = 0x48
		//	ExpectedProcessorStatusValue: 0b10000000, // Carry remains set
		//},
		{
			Name:                         "ADC Zero Page",
			AccumolatorSetup:             0x42,
			MemorySetup:                  []any{c.ADC_Z, 0x05, 0x05},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b10000000},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x47,
			ExpectMemory:                 []any{0x05: 0x05},
			ExpectedProcessorStatusValue: 0b00000000,
		},
		{
			Name:                         "ADC Zero Page,X with Carry",
			AccumolatorSetup:             0x42,
			RegisterXSetup:               0x05,
			MemorySetup:                  []any{c.ADC_ZX, 0x05, 0x05},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b10000000},
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x47,
			ExpectMemory:                 []any{0x0A: 0x05},
			ExpectedProcessorStatusValue: 0b00000000,
		}, //{
		//	Name: "ADC Absolute",
		//	AccumolatorSetup: 0x42,
		//	MemorySetup: []any{c.ADC_A, 0x12, 0x34},
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x47,
		//	ExpectMemory: []any{0x1234: 0x05},
		//	ExpectedProcessorStatusValue: 0b00000000,
		//},
		//{
		//	Name: "ADC Absolute,X with Carry",
		//	AccumolatorSetup: 0x42,
		//	RegisterXSetup: 0x05,
		//	MemorySetup: []any{c.ADC_AX, 0x12, 0x34},
		//	ProcessorStatusValueSetup: 0b10000000,
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x47,
		//	ExpectMemory: []any{0x1239: 0x05},
		//	ExpectedProcessorStatusValue: 0b10000000,
		//},
		//{
		//	Name: "ADC Absolute,Y with Carry",
		//	AccumolatorSetup: 0x42,
		//	RegisterYSetup: 0x05,
		//	MemorySetup: []any{c.ADC_AY, 0x12, 0x34},
		//	ProcessorStatusValueSetup: 0b10000000,
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x47,
		//	ExpectMemory: []any{0x1239: 0x05},
		//	ExpectedProcessorStatusValue: 0b10000000,
		//},
		//{
		//	Name: "ADC (Indirect,X) with Carry",
		//	AccumolatorSetup: 0x42,
		//	RegisterXSetup: 0x05,
		//	MemorySetup: []any{c.ADC_IX, 0x10, 0x20, 0x05},
		//	ProcessorStatusValueSetup: 0b10000000,
		//	ExpectToAdvancedCycles: 6,
		//	ExpectAccumulatorValue: 0x47,
		//	ExpectMemory: []any{0x0025: 0x05},
		//	ExpectedProcessorStatusValue: 0b10000000,
		//},
		//{
		//	Name: "ADC (Indirect),Y with Carry",
		//	AccumolatorSetup: 0x42,
		//	RegisterYSetup: 0x05,
		//	MemorySetup: []any{c.ADC_IY, 0x10, 0x20, 0x05},
		//	ProcessorStatusValueSetup: 0b10000000,
		//	ExpectToAdvancedCycles: 5,
		//	ExpectAccumulatorValue: 0x47,
		//	ExpectMemory: []any{0x0025: 0x05},
		//	ExpectedProcessorStatusValue: 0b10000000,
		//},
		//{
		//	Name: "ADC Immediate with Overflow",
		//	AccumolatorSetup: 0x80,
		//	MemorySetup: []any{c.ADC_IMM, 0x40},
		//	ExpectToAdvancedCycles: 2,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000, // Carry and Overflow set
		//},
		{
			Name:                         "ADC Zero Page with Overflow",
			AccumolatorSetup:             0x80,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0xa0},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b00000000},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x20,
			ExpectedProcessorStatusValue: c.OverflowFlag | c.CarryFlag,
		}, //{
		//	Name: "ADC Absolute with Overflow",
		//	AccumolatorSetup: 0x80,
		//	MemorySetup: []any{c.ADC_A, 0x12, 0x34},
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectMemory: []any{0x1234: 0x40},
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000,
		//},
		{
			Name:                         "ADC Zero Page,X with Overflow",
			AccumolatorSetup:             0x80,
			RegisterXSetup:               0x05,
			MemorySetup:                  []any{c.ADC_ZX, 0x00, 0xA0},
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: 0b00000000},
			ExpectToAdvancedCycles:       4,
			ExpectAccumulatorValue:       0x20,
			ExpectedProcessorStatusValue: c.OverflowFlag | c.CarryFlag,
		}, //{
		//	Name: "ADC Absolute,X with Overflow",
		//	AccumolatorSetup: 0x80,
		//	RegisterXSetup: 0x05,
		//	MemorySetup: []any{c.ADC_AX, 0x12, 0x34},
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectMemory: []any{0x1239: 0x40},
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000,
		//},
		//{
		//	Name: "ADC Absolute,Y with Overflow",
		//	AccumolatorSetup: 0x80,
		//	RegisterYSetup: 0x05,
		//	MemorySetup: []any{c.ADC_AY, 0x12, 0x34},
		//	ExpectToAdvancedCycles: 4,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectMemory: []any{0x1239: 0x40},
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000,
		//},
		//{
		//	Name: "ADC (Indirect,X) with Overflow",
		//	AccumolatorSetup: 0x80,
		//	RegisterXSetup: 0x05,
		//	MemorySetup: []any{c.ADC_IX, 0x10, 0x20, 0x40},
		//	ExpectToAdvancedCycles: 6,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectMemory: []any{0x0025: 0x40},
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000,
		//},
		//{
		//	Name: "ADC (Indirect),Y with Overflow",
		//	AccumolatorSetup: 0x80,
		//	RegisterYSetup: 0x05,
		//	MemorySetup: []any{c.ADC_IY, 0x10, 0x20, 0x40},
		//	ExpectToAdvancedCycles: 5,
		//	ExpectAccumulatorValue: 0x20,
		//	ExpectMemory: []any{0x0025: 0x40},
		//	ExpectedProcessorStatusValue: 0b10000000 | 0b01000000,
		//},
		{
			Name:                         "ADC Zero Page - Simple Addition",
			AccumolatorSetup:             0x10,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x20},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x30,
			ExpectedProcessorStatusValue: 0,
		},
		{
			Name:                         "ADC Zero Page - Uses Carry In",
			AccumolatorSetup:             0x10,
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: c.Word(c.CarryFlag)},
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x20},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x31,
			ExpectedProcessorStatusValue: 0,
		},
		{
			Name:                         "ADC Zero Page - Carry Out and Zero Result",
			AccumolatorSetup:             0xFF,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x01},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x00,
			ExpectedProcessorStatusValue: c.CarryFlag | c.ZeroFlag,
		},
		{
			Name:                         "ADC Zero Page - Carry In Causes Carry Out and Zero Result",
			AccumolatorSetup:             0xFF,
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: c.Word(c.CarryFlag)},
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x00},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x00,
			ExpectedProcessorStatusValue: c.CarryFlag | c.ZeroFlag,
		},
		{
			Name:                         "ADC Zero Page - Positive Signed Overflow",
			AccumolatorSetup:             0x7F,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x01},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x80,
			ExpectedProcessorStatusValue: c.NegativeFlag | c.OverflowFlag,
		},
		{
			Name:                         "ADC Zero Page - Negative Signed Overflow",
			AccumolatorSetup:             0x80,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x80},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x00,
			ExpectedProcessorStatusValue: c.CarryFlag | c.ZeroFlag | c.OverflowFlag,
		},
		{
			Name:                         "ADC Zero Page - Overflow and Carry Without Negative",
			AccumolatorSetup:             0x80,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0xA0},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x20,
			ExpectedProcessorStatusValue: c.CarryFlag | c.OverflowFlag,
		},
		{
			Name:                         "ADC Zero Page - Negative Result Without Overflow",
			AccumolatorSetup:             0x80,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x01},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x81,
			ExpectedProcessorStatusValue: c.NegativeFlag,
		},
		{
			Name:                         "ADC Zero Page - Different Signs No Overflow",
			AccumolatorSetup:             0xFF,
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x02},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x01,
			ExpectedProcessorStatusValue: c.CarryFlag,
		},
		{
			Name:                         "ADC Zero Page - Carry In Creates Signed Overflow",
			AccumolatorSetup:             0x7F,
			ProcessorStatusValueSetup:    c.ProcessorStatus{Status: c.Word(c.CarryFlag)},
			MemorySetup:                  []any{c.ADC_Z, 0x00, 0x00},
			ExpectToAdvancedCycles:       3,
			ExpectAccumulatorValue:       0x80,
			ExpectedProcessorStatusValue: c.NegativeFlag | c.OverflowFlag,
		},
	}

	t.Logf("All tests for CLC")
	for idx, testData := range data {
		testData.Name = strconv.Itoa(idx+1) + "_" + testData.Name

		cpu := c.NewSixFiveOTwo(ut.NewTestCpuLogger(t))
		tm := ut.DefaultTestMemory(t)

		testData.Run(t, cpu, tm)
	}

}
