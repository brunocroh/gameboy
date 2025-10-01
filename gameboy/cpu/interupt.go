package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

const IE_ADDRESS = uint16(0xFFFF) // Interuptor Enabled memory location
const IF_ADDRESS = uint16(0xFF0F) // Interuptor Flags memory location

type interupt struct {
	IME bool
	mmu *mmu.MemoryManagementUnit
}

func interuptNew(mmu *mmu.MemoryManagementUnit) *interupt {
	return &interupt{
		mmu: mmu,
	}
}

func (m *interupt) Init() {
	m.IME = false
}

func (m *interupt) HandleInterupt() uint16 {
	if !m.IME {
		return 0
	}

	inte := m.mmu.RB(IE_ADDRESS)
	intf := m.mmu.RB(IF_ADDRESS)

	triggered := inte & intf

	fmt.Println("inte:", inte)
	fmt.Println("intf:", intf)

	if triggered == 0 {
		return 0
	}

	return 4
}
