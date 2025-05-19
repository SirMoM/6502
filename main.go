package main

import (
	"fmt"
	"noah-ruben.com/6502/computer"
	"noah-ruben.com/6502/programs"
)

func main() {
	println("6502")

	logger := computer.SetupLogging()

	cpu := computer.NewSixFiveOTwo(&logger)
	mem := computer.Memory16K{}

	_ = mem.Init()

	cpu.Reset(&mem)
	_ = programs.MiniProg.CopyToMemory(cpu.ProgramCounter, &mem)

	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(2)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(4)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(8)
	fmt.Println(cpu)

	_ = logger.Close()
}
