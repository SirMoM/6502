package main

import (
	"fmt"
)

func main() {
	println("6502")
	cpu := SixFiveOTwo{}
	mem := Memory{}
	mem.Init()
	cpu.Reset(&mem)

	fmt.Println(cpu)
	fmt.Println("\n###############################")
	fmt.Println(mem)
	fmt.Println(cpu)
	cpu.Execute(5, &mem)
	fmt.Println(cpu)

}
