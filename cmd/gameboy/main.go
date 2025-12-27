package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/brunocroh/gameboy/gameboy"
)

func main() {

	romPtr := flag.String("rom", "", "rom to execute")
	singleStepPtr := flag.Bool("single-step", false, "enable single step execution")

	flag.Parse()

	gb := gameboy.New()
	gb.Init(*romPtr)

	reader := bufio.NewReader(os.Stdin)

	for {
		if *singleStepPtr {
			_, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("fail to read", err)
				return
			}

		}

		// time.Sleep(100 * time.Millisecond)

		gb.Update()
	}
}
