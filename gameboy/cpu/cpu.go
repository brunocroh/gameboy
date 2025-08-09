package cpu

import "fmt"

type Cpu struct {
}

func Cycle() {

	fetch()
	decode()
	instruction()
}

func fetch() {
	fmt.Println("fetch")

}

func decode() {
	fmt.Println("fetch")
}

func instruction() {
	fmt.Println("instruction")
}
