package main

import "fmt"

func main() {
	println("6502")

	cpu := SixFiveOTwo{}
	mem := Memory{}
	mem.Init()
	cpu.Reset(&mem)

	cpu.Execute(1, &mem, true)
	cpu.Execute(1, &mem, true)
	fmt.Println(cpu)
}
