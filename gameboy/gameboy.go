package gameboy

import (
	"fmt"

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
	m.mmu.Init()
	m.cpu = cpu.New(m.mmu)
}

func (m *GameBoy) Debug() {
	fmt.Println(m.mmu.Dump())
}
