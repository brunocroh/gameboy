package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	ins *instructions

	mmu *mmu.MemoryManagementUnit

	pc uint16
	sp uint16

	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8

	f uint8
}

func New(mmu *mmu.MemoryManagementUnit) *CPU {
	i := instructionsNew()
	return &CPU{
		mmu: mmu,
		ins: i,
	}
}

func (m *CPU) Init() {
	m.pc = 0x0100
	m.sp = 0xFFFE
	m.a = 0x01
	m.f = 0xB0 //check later the initial value and how to handle flags
	m.b = 0x00
	m.c = 0x13
	m.d = 0x00
	m.e = 0xD8
	m.h = 0x01
	m.l = 0x4D
}

func (m *CPU) Cycle() {
	opcode := m.fetchOpcode()
	m.execInstruction(opcode)
}

func (m *CPU) fetchOpcode() byte {
	pc := m.pc
	m.pc += 1

	return m.mmu.RB(pc)
}

func (m *CPU) execInstruction(opcode byte) {
	var ticks uint32
	switch opcode {
	case 0xC3:
		ticks = m.ins.jpAddr(m, m.mmu.RW(m.pc))
	default:
		fmt.Printf("opcode (0x%x) not implemented\n", opcode)
	}

	if ticks != 0 {
		m.doCycle(ticks)
	}
}

func (m *CPU) doCycle(ticks uint32) uint32 {
	return m.mmu.DoCycle(ticks * 4)
}
