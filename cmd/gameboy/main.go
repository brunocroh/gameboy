package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/brunocroh/gameboy/gameboy"
)

func main() {
	fmt.Println("cli/desktop version of gameboy emulator")
	romPath := os.Args[1:]

	gb := gameboy.New()
	gb.Init(romPath[0])

	reader := bufio.NewReader(os.Stdin)

	for {
		_, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("fail to read", err)
			return
		}

		gb.Update()
	}
}
