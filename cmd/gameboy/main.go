package main

import (
	"fmt"
	"os"

	"github.com/brunocroh/gameboy/gameboy"
)

func main() {
	fmt.Println("cli/desktop version of gameboy emulator")
	romPath := os.Args[1:]

	gb := gameboy.New()
	gb.Init(romPath[0])
	gb.Debug()

	for i := 0; i < 10; i++ {
		gb.Update()
	}

	gb.Debug()
}
