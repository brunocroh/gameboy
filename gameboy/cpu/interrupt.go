package cpu

import (
	"github.com/brunocroh/gameboy/gameboy/mmu"
)

const IE_ADDRESS = uint16(0xFFFF) // Interuptor Enabled memory location
const IF_ADDRESS = uint16(0xFF0F) // Interuptor Flags memory location

type interrupt struct {
	IME bool
	mmu *mmu.MemoryManagementUnit
}

func interruptNew(mmu *mmu.MemoryManagementUnit) *interrupt {
	return &interrupt{
		mmu: mmu,
	}
}

func (m *interrupt) Init() {
	m.IME = false
}

func (m *interrupt) handleInterrupt() uint32 {
	if !m.IME {
		return 0
	}

	inte := m.mmu.RB(IE_ADDRESS)
	intf := m.mmu.RB(IF_ADDRESS)

	triggered := inte & intf

	if triggered == 0 {
		return 0
	}

	return 4
}
