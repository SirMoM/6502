package main

import (
	"fmt"
	"noah-ruben.com/6502/computer"
	"noah-ruben.com/6502/programs"
)

func main() {
	println("6502")

	cpu := computer.SixFiveOTwo{}
	mem := computer.Memory{}
	mem.Init()
	_ = programs.MiniProg.CopyToMemory(cpu.ProgramCounter, &mem)
	cpu.Reset(&mem)

	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(2)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(4)
	cpu.Execute(1, &mem, true)
	cpu.AssertCycle(8)
	fmt.Println(cpu)
}
