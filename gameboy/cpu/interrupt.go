package cpu

import (
	"github.com/brunocroh/gameboy/gameboy/mmu"
)

const IE_ADDRESS = uint16(0xFFFF) // Interuptor Enabled memory location
const IF_ADDRESS = uint16(0xFF0F) // Interuptor Flags memory location

type interrupt struct {
	IME  bool
	EI   uint8
	Halt uint8
	mmu  mmu.MemoryManagementUnit
}

func interruptNew(mmu mmu.MemoryManagementUnit) *interrupt {
	return &interrupt{
		mmu: mmu,
	}
}

func (m *interrupt) Init() {
	m.IME = false
}

func (m *interrupt) handleInterrupt(cpu *CPU) uint32 {
	enabled := m.mmu.RB(IE_ADDRESS)
	requested := m.mmu.RB(IF_ADDRESS)

	triggered := enabled & requested & 0x1F

	if triggered == 0 {
		return 0
	}

	m.Halt = 0
	if !m.IME {
		return 0
	}

	var mask uint8
	var vector uint16

	switch {
	case triggered&0x01 != 0:
		mask = 0x01
		vector = 0x40
	case triggered&0x02 != 0:
		mask = 0x02
		vector = 0x48
	case triggered&0x04 != 0:
		mask = 0x04
		vector = 0x50
	case triggered&0x08 != 0:
		mask = 0x08
		vector = 0x58
	case triggered&0x10 != 0:
		mask = 0x10
		vector = 0x60
	}

	m.IME = false

	requested &^= mask
	m.mmu.WB(IF_ADDRESS, requested)

	// push to stack
	cpu.SP--
	cpu.mmu.WB(cpu.SP, uint8(cpu.PC>>8))
	cpu.SP--
	cpu.mmu.WB(cpu.SP, uint8(cpu.PC))

	cpu.PC = vector

	return 5
}

func (m *interrupt) updateIME() {
	switch m.EI {
	case 2:
		m.EI = 1
	case 1:
		m.IME = true
		m.EI = 0
	}
}

func (m *interrupt) Enable() uint32 {
	m.EI = 2
	return 1
}

func (m *interrupt) Disable() uint32 {
	m.IME = false
	m.EI = 0
	return 1
}
