package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	mmu *mmu.MemoryManagementUnit

	ins      *instructions
	register *register

	pc uint16
	sp uint16
}

func New(mmu *mmu.MemoryManagementUnit) *CPU {
	i := instructionsNew()
	r := registerNew()
	return &CPU{
		mmu:      mmu,
		ins:      i,
		register: r,
	}
}

func (m *CPU) Init() {
	m.pc = 0x0100
	m.sp = 0xFFFE
	m.register.Init()
}

func (m *CPU) Cycle() {
	opcode := m.fetchOpcode()
	m.doCycle(1)
	m.execInstruction(opcode)
}

func (m *CPU) fetchOpcode() byte {
	return m.rb(m.pc)
}

func (m *CPU) execInstruction(opcode byte) {
	var ticks uint32
	switch opcode {
	case 0x06:
		ticks = m.ins.ldBAddress(m, m.pc)
	case 0xC3:
		ticks = m.ins.jpAddr(m, m.pc)
	case 0x41:
		ticks = m.ins.ldBC(m)
	default:
		fmt.Printf("opcode (0x%x) not implemented\n", opcode)
	}

	if ticks != 0 {
		m.doCycle(ticks)
	}
}

func (m *CPU) rb(addr uint16) byte {
	value := m.mmu.RB(addr)
	m.pc += 1
	return value
}

func (m *CPU) rw(addr uint16) uint16 {
	value := m.mmu.RW(addr)
	m.pc += 2
	return value
}

func (m *CPU) doCycle(ticks uint32) uint32 {
	return m.mmu.DoCycle(ticks * 4)
}
