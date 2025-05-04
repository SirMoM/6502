package main

import "fmt"

func main() {
	println("6502")
	cpu := SixFiveOTwo{}
	mem := Memory{}
	mem.Init()
	cpu.Reset(&mem)

	cpu.Execute(3, &mem)
	fmt.Println(cpu)
	cpu.Execute(1, &mem)
}
