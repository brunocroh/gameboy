package gameboy

import (
	"fmt"
	"os"

	"github.com/brunocroh/gameboy/gameboy/cpu"
	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type GameBoy struct {
	cpu *cpu.CPU
	mmu *mmu.MemoryManagementUnit
}

func New() *GameBoy {
	return &GameBoy{}
}

func (m *GameBoy) Init() {
	m.mmu = mmu.New()
	m.cpu = cpu.New(m.mmu)
	m.cpu.Init()
}

func (m *GameBoy) Update() {
	m.cpu.Cycle()
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
