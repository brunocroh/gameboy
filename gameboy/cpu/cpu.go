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
	var ticks uint32 = 1
	switch opcode {
	case 0x01:
		ticks = m.ins.ld_rr_nn(m)
	case 0x02:
		ticks = m.ins.ld_BC_A(m)
	case 0x04:
		ticks = m.ins.inc_r(m)
	case 0x05:
		ticks = m.ins.dec_r(m)
	case 0x06:
		ticks = m.ins.ld_r_n(m)
	case 0x08:
		ticks = m.ins.ld_nn_sp(m)
	case 0x0A:
		ticks = m.ins.ld_A_BC(m)
	case 0x12:
		ticks = m.ins.ld_DE_A(m)
	case 0x1A:
		ticks = m.ins.ld_A_DE(m)
	case 0x22:
		ticks = m.ins.ld_HLi_A(m)
	case 0x2A:
		ticks = m.ins.ld_A_HLi(m)
	case 0x32:
		ticks = m.ins.ld_HLd_A(m)
	case 0x34:
		ticks = m.ins.inc_HL(m)
	case 0x35:
		ticks = m.ins.dec_HL(m)
	case 0x36:
		ticks = m.ins.ld_HL_n(m)
	case 0x3A:
		ticks = m.ins.ld_A_HLd(m)
	case 0x41:
		ticks = m.ins.ld_rr(m)
	case 0x46:
		ticks = m.ins.ld_r_HL(m)
	case 0x70:
		ticks = m.ins.ld_HL_r(m)
	case 0x80:
		ticks = m.ins.add_r(m)
	case 0x86:
		ticks = m.ins.add_HL(m)
	case 0x88:
		ticks = m.ins.adc_r(m)
	case 0x8E:
		ticks = m.ins.adc_HL(m)
	case 0x90:
		ticks = m.ins.sub_r(m)
	case 0x96:
		ticks = m.ins.sub_HL(m)
	case 0x98:
		ticks = m.ins.sbc_r(m)
	case 0x9E:
		ticks = m.ins.sbc_HL(m)
	case 0xB8:
		ticks = m.ins.cp_r(m)
	case 0xBE:
		ticks = m.ins.cp_HL(m)
	case 0xC1:
		ticks = m.ins.ld_pop_rr(m)
	case 0xC3:
		ticks = m.ins.jp_nn(m)
	case 0xC5:
		ticks = m.ins.ld_push_rr(m)
	case 0xC6:
		ticks = m.ins.add_n(m)
	case 0xCE:
		ticks = m.ins.adc_n(m)
	case 0xD6:
		ticks = m.ins.sub_n(m)
	case 0xDE:
		ticks = m.ins.sbc_n(m)
	case 0xE0:
		ticks = m.ins.ldh_n_A(m)
	case 0xE2:
		ticks = m.ins.ldh_C_A(m)
	case 0xEA:
		ticks = m.ins.ld_nn_A(m)
	case 0xF0:
		ticks = m.ins.ldh_A_n(m)
	case 0xF2:
		ticks = m.ins.ldh_A_C(m)
	case 0xF8:
		ticks = m.ins.ld_HL_spe(m)
	case 0xF9:
		ticks = m.ins.ld_sp_HL(m)
	case 0xFA:
		ticks = m.ins.ld_A_nn(m)
	case 0xFE:
		ticks = m.ins.cp_n(m)
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
