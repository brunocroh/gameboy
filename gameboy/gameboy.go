package gameboy

import "fmt"

type Gameboy struct {
	memory [0xFFFF]uint16 //64KB
}

func Init() {
	fmt.Println("emulator entrypoint")
}
