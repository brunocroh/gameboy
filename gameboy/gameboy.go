package gameboy

import (
	"fmt"
	"os"

	"github.com/brunocroh/gameboy/gameboy/cpu"
	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type GameBoy struct {
	cpu *cpu.CPU
	mmu mmu.MemoryManagementUnit
}

func New() *GameBoy {
	return &GameBoy{}
}

func (m *GameBoy) Init(filePath string) {
	m.mmu = mmu.NewMemoryManagementUnitSimple()
	rom, err := LoadROM(filePath)

	if err != nil {
		fmt.Println("FAIL TO LOAD ROM")
	}

	m.mmu.Init(rom)
	m.cpu = cpu.New(m.mmu)
	m.cpu.Init()
}

func (m *GameBoy) Update() {
	m.cpu.Cycle()
	m.Debug()
}

func (m *GameBoy) Debug() {
	fmt.Println("======== DEBUG =========")
	fmt.Println(m.mmu.Dump())
	fmt.Println("========================")
}

func LoadROM(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return data, nil
}
