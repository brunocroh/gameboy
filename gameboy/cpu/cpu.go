package cpu

import (
	"fmt"

	"github.com/brunocroh/gameboy/gameboy/mmu"
)

type CPU struct {
	mmu *mmu.MemoryManagementUnit

	ins      *instructions
	register *register

	PC uint16
	SP uint16
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
	m.PC = 0x0100
	m.SP = 0xFFFE
	m.register.Init()
}

func (m *CPU) Cycle() {
	opcode := m.fetchOpcode()
	m.doCycle(1)
	m.execInstruction(opcode)
}

func (m *CPU) popPC() uint16 {
	pc := m.PC
	m.PC += 1
	return pc
}

func (m *CPU) fetchOpcode() byte {
	return m.mmu.RB(m.popPC())
}

func (m *CPU) execInstruction(opcode byte) {
	var ticks uint32
	switch opcode {
	case 0x06:
		ticks = m.ins.ld_r_n(m, m.PC)
	case 0xC3:
		ticks = m.ins.jp_nn(m, m.PC)
	case 0x41:
		ticks = m.ins.ld_rr(m)
	default:
		fmt.Printf("opcode (0x%x) not implemented\n", opcode)
	}

	if ticks != 0 {
		m.doCycle(ticks)
	}
}

func (m *CPU) rw(addr uint16) uint16 {
	value := m.mmu.RW(addr)
	m.PC += 2
	return value
}

func (m *CPU) doCycle(ticks uint32) uint32 {
	return m.mmu.DoCycle(ticks * 4)
}
