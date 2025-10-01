package main

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy"
)

func main() {
	fmt.Println("cli/desktop version of gameboy emulator")

	gb := gameboy.New()
	gb.Init()
	gb.Debug()

	for i := 0; i < 10; i++ {
		gb.Update()
	}

	gb.Debug()
}
